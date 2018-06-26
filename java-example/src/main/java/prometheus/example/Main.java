package prometheus.example;

import io.dropwizard.Application;
import io.dropwizard.Configuration;
import io.dropwizard.setup.Environment;
import prometheus.example.resources.HelloResource;

import java.util.logging.Logger;

public class Main {
    private static final Logger LOG = Logger.getLogger(Main.class.getName());

    private static class WebConfig extends Configuration {}

    private static class WebServer extends Application<WebConfig> {
        @Override
        public void run(WebConfig config, Environment env) {
            env.jersey().register(HelloResource.class);
        }
    }

    public static void main(String[] args) throws Exception {
        LOG.info("Starting up the server on port 8080...");
        new WebServer().run(args);
    }
}
