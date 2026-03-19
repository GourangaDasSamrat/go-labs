const port = 8000;

const server = Bun.serve({
  port,
  async fetch(req: Request): Promise<Response> {
    const url = new URL(req.url);
    const { pathname } = url;

    // Helper to parse body (JSON or form)
    async function parseBody(req: Request): Promise<any> {
      const contentType = req.headers.get("content-type") || "";

      if (contentType.includes("application/json")) {
        return await req.json();
      }

      if (contentType.includes("application/x-www-form-urlencoded")) {
        const formData = await req.formData();
        return Object.fromEntries(formData.entries());
      }

      return null;
    }

    // Routes
    if (req.method === "GET" && pathname === "/") {
      return new Response("Welcome to Gouranga's server", {
        status: 200,
      });
    }

    if (req.method === "GET" && pathname === "/get") {
      return Response.json(
        { message: "Hello from Gouranga" },
        { status: 200 }
      );
    }

    if (req.method === "POST" && pathname === "/post") {
      const body = await parseBody(req);

      return new Response(JSON.stringify(body), {
        status: 200,
        headers: { "Content-Type": "application/json" },
      });
    }

    if (req.method === "POST" && pathname === "/postform") {
      const body = await parseBody(req);

      return new Response(JSON.stringify(body), {
        status: 200,
      });
    }

    // 404 fallback
    return new Response("Not Found", { status: 404 });
  },
});

console.log(`Server running at http://localhost:${server.port}`);
