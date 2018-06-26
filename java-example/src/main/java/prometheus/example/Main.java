package prometheus.example;

import io.dropwizard.Application;
import io.dropwizard.Configuration;
import io.dropwizard.setup.Environment;

public class Main {
    private static class WebConfig extends Configuration {}

    private static class WebServer extends Application<WebConfig> {
        @Override
        public void run(WebConfig config, Environment env) {

        }
    }

    public static void main(String[] args) throws Exception {
        new WebServer().run(args);
    }
}
