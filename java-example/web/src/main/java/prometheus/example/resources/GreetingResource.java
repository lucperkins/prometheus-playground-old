package prometheus.example.resources;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.atomic.AtomicLong;

@RestController
public class GreetingResource {
    private static final String template = "Hello, %s!";
    private static final AtomicLong counter = new AtomicLong(0);
    private static final Counter greetingCounter = Metrics.counter("greetings");

    @RequestMapping("/greeting")
    public Greeting greet(@RequestParam(value = "name", defaultValue = "world") String name) {
        greetingCounter.increment();

        return new Greeting(counter.incrementAndGet(),
                String.format(template, name));
    }

    private static class Greeting {
        private final long id;
        private final String content;

        public Greeting(long id, String content) {
            this.id = id;
            this.content = content;
        }

        public long getId() {
            return id;
        }

        public String getContent() {
            return content;
        }
    }


}
