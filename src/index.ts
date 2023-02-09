import http from "http";

const server = http.createServer(async (request, response) => {
    try {
        const buffers = [];

        for await (const chunk of request) {
            buffers.push(chunk);
        }

        const data = Buffer.concat(buffers).toString();

        console.log(`${request.method} -> ${decodeURI(request.url ?? "")}`, request.rawHeaders, data);

        response.writeHead(200);
        response.end();
    } catch (error) {
        console.log("Unexpected error occurred!", error);
        response.writeHead(500);
        response.end();
    }
});

const port = process.env.PORT ?? 5000;
server.listen(port, () => {
    console.log(`Server listens on port ${port}`);
});
