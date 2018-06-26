package prometheus.example.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.core.Response;

@Path("/")
public class HelloResource {
    @GET
    public Response helloWorld() {
        return Response
                .status(Response.Status.OK)
                .entity("Hello world")
                .build();
    }
}
