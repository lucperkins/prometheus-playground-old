package prometheus.example.health;

import com.codahale.metrics.health.HealthCheck;

public class Health extends HealthCheck {
    @Override
    public Result check() {
        return Result.healthy("Everything is turning up Millhouse");
    }
}
