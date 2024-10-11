#![forbid(unsafe_code)]

use std::{
    env,
    fs::{self, File},
    io::Write,
    path::PathBuf,
};

use axum::{
    body::Body,
    extract::{Multipart, Request},
    http::{header::CONTENT_TYPE, StatusCode},
    response::{IntoResponse, Response},
    routing::any,
    RequestExt, Router,
};

use std::io::Result;

use uuid::Uuid;

const FILE_FOLDER_PATH: &str = "files";

#[tokio::main]
async fn main() {
    let folder_path = env::current_dir()
        .expect("CWD should be accessible")
        .join(FILE_FOLDER_PATH);

    if !folder_path.exists() {
        fs::create_dir(folder_path).expect("Should have sufficient FS permissions for CWD");
    }

    let app = Router::new()
        .route("/", any(handler))
        .route("/*_", any(handler));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app)
        .await
        .expect("Server should be able to start");
}

async fn handler(req: Request<Body>) -> Response {
    println!("{0} -> {1}", req.method().as_str(), req.uri().to_string());

    let content_type = req
        .headers()
        .get(CONTENT_TYPE)
        .and_then(|x| x.to_str().ok())
        .unwrap_or_default();

    if content_type.starts_with("multipart/form-data") {
        match req.extract::<Multipart, _>().await {
            Ok(multi) => handle_multipart(multi).await,
            Err(e) => println!("Could not resolve multipart request! {0}", e),
        }
    } else {
        match req.extract::<String, _>().await {
            Ok(body) => println!("{}", body),
            Err(e) => println!("Could not resolve request body! {0}", e),
        }
    }

    (StatusCode::OK).into_response()
}

async fn handle_multipart(mut mult: Multipart) {
    while let Ok(Some(field)) = mult.next_field().await {
        let name = field.name().unwrap_or_default().to_owned();
        let content_type = field.content_type().unwrap_or_default().to_owned();

        let Ok(data) = field.text().await else {
            println!("Cloud not resolve data of field {0}!", name);
            return;
        };

        if content_type.is_empty() {
            println!("{0} = {1}", name, data);
        } else {
            match save_new_file(data) {
                Ok(path) => println!(
                    "{0} = Saved to: {1}",
                    name,
                    path.to_str().unwrap_or_default()
                ),
                Err(e) => println!("Cloud not save data for field {0}! {1}", name, e),
            }
        }
    }
}

fn save_new_file(data: String) -> Result<PathBuf> {
    let path = env::current_dir()?
        .join(FILE_FOLDER_PATH)
        .join(Uuid::new_v4().to_string());

    if let Err(e) = File::create_new(path.clone())?.write_all(data.as_bytes()) {
        fs::remove_file(path)?;
        return Err(e);
    }

    Ok(path)
}
