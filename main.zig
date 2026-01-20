const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();
    
    var stdout_buffer: [1024]u8 = undefined;
    var stdout_writer = std.fs.File.stdout().writer(&stdout_buffer);
    const stdout = &stdout_writer.interface;
    
    // Get the number of rounds (cross-platform).
    const strRounds = std.process.getEnvVarOwned(allocator, "rounds") catch |err| {
        if (err == error.EnvironmentVariableNotFound) {
            try stdout.print("*** The 'rounds' environment variable is not set\n", .{});
        } else {
            try stdout.print("*** Error reading 'rounds' environment variable\n", .{});
        }
        try stdout.flush();
        std.process.exit(1);
    };
    defer allocator.free(strRounds);
    
    const rounds = std.fmt.parseInt(i64, strRounds, 10) catch {
        try stdout.print("*** The 'rounds' environment variable must be an integer, saw: {s}\n", .{strRounds});
        try stdout.flush();
        std.process.exit(1);
    };
    
    var sum: f64 = 0.0;
    var flip: f64 = -1.0;
    var pi: f64 = undefined;
    var t_start: i64 = undefined;
    var t_stop: i64 = undefined;
    
    // Prime the caches.
    var ix: i64 = 1;
    while (ix <= 3) : (ix += 1) {
        flip *= -1.0;
        sum += flip / @as(f64, @floatFromInt(ix + ix - 1));
    }
    sum = 0.0;
    flip = -1.0;
    
    // Timed test.
    t_start = std.time.milliTimestamp();
    ix = 1;
    while (ix <= rounds) : (ix += 1) {
        flip *= -1.0;
        sum += flip / @as(f64, @floatFromInt(ix + ix - 1));
    }
    pi = sum * 4.0;
    t_stop = std.time.milliTimestamp();
    
    // Report.
    const t_delta_secs = @as(f64, @floatFromInt(t_stop - t_start)) / 1000.0;
    try stdout.print("Zig,{d},{d:.3},{d:.40}\n", .{rounds, t_delta_secs, pi});
    try stdout.flush();
}
