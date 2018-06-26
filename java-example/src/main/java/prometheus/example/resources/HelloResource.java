package prometheus.example.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.core.Response;

@Path("/")
public class HelloResource {
    @GET
    @Path("/{name}")
    public Response helloWorld(@PathParam("name") final String name) {
        String greeting = String.format("Hello, %s!", name);

        return Response
                .status(Response.Status.OK)
                .entity(greeting)
                .build();
    }
}
