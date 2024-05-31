const std = @import("std");
const expect = std.testing.expect;

test "always succeeds" {
    try expect(true);
}

test "demo1"{
    var a:u32 = 1;
    const b = 2;
    const iv1 = @as(i8, 5);
    var iv2 = @as(i32, 5000);
    iv2 += iv1;
    std.debug.print("i:{d} \n", .{iv1+iv2});

    a += 2;

    if (a > b) {
        try expect(true);
    }else {
        try expect(false);
    }
}

test "demo list" {
    const a = [5]u8{'h','e','l','l','o'};
    const b = [_]u8{'w','o','r','l','d'};

    std.debug.print(",,,,,,,,,,,,,,, {d} {d}\n", .{a.len, b.len});
}

test "demo if"{
    const a = true;
    var x:u16 = 0;
    x += if (a) 1 else 2;
    std.debug.print(" -       -  {d}", .{x});
}

test "demo while" {
    var sum:u8 = 0;
    var i:u8 = 1;
    while (i<=10): (i+=1) {
        sum += i;
    }
    std.debug.print("-- {d} sum", .{sum});
}

test "demo for" {
    const strs = [_]u8{'a','b','c'};
    for (strs,0..) |character, index| {
        std.debug.print("--: {} {}", .{character,index});
    }
}

test "demo defer"{
    var x:i16 = 5;
    {
        defer x += 2;
        std.debug.print("--- v1: {}", .{x}); // 5
    }
    std.debug.print("--- v2: {}", .{x}); // 7
}