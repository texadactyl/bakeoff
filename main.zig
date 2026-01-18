const std = @import("std");

pub fn main() !void {
    const rounds: i64 = 3_000_000_000;
    
    var stdout_buffer: [1024]u8 = undefined;
    var stdout_writer = std.fs.File.stdout().writer(&stdout_buffer);
    const stdout = &stdout_writer.interface;
    
    var sum: f64 = 0.0;
    var flip: f64 = -1.0;
    var pi: f64 = undefined;
    var ix: i64 = undefined;
    
    // Prime the caches.
    ix = 1;
    while (ix <= 3) : (ix += 1) {
        flip *= -1.0;
        sum += flip / @as(f64, @floatFromInt(ix + ix - 1));
    }
    sum = 0.0;
    flip = -1.0;
    
    // Timed test.
    var timer = try std.time.Timer.start();
    
    ix = 1;
    while (ix <= rounds) : (ix += 1) {
        flip *= -1.0;
        sum += flip / @as(f64, @floatFromInt(ix + ix - 1));
    }
    
    pi = sum * 4.0;
    const elapsed_ns = timer.read();
    
    // Report.
    const elapsed_secs = @as(f64, @floatFromInt(elapsed_ns)) / 1e9;
    try stdout.print("Zig,{d},{d:.3},{d:.16}\n", .{rounds, elapsed_secs, pi});
    try stdout.flush();
}
