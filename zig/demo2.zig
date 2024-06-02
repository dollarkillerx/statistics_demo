const std = @import("std");
const hashmap = std.hash_map.StringHashMap;

const User = struct {
    email: []const u8,
    password: []const u8,
    amount: f64,
    };

const Bank = struct {
    users: hashmap(User),

    pub fn init(allocator: *std.mem.Allocator) !Bank {
        const users = hashmap(User).init(allocator);
        return Bank{ .users = users };
    }

    pub fn addUser(self: *Bank, email: []const u8, password: []const u8) !void {
        const user = User{
            .email = email,
            .password = password,
            .amount = 0.0,
        };
        self.users.put(email, user) catch return errors.UserExists;
    }

    pub fn deposit(self: *Bank, email: []const u8, amount: f64) !void {
        const user = self.users.get(email) orelse return errors.UserNotFound;
        user.amount += amount;
    }

    pub fn withdraw(self: *Bank, email: []const u8, amount: f64) !void {
        const user = self.users.get(email) orelse return errors.UserNotFound;
        if (user.amount < amount) return errors.InsufficientFunds;
        user.amount -= amount;
    }

    pub fn transfer(self: *Bank, from_email: []const u8, to_email: []const u8, amount: f64) !void {
        const from_user = self.users.get(from_email) orelse return errors.UserNotFound;
        const to_user = self.users.get(to_email) orelse return errors.UserNotFound;
        if (from_user.amount < amount) return errors.InsufficientFunds;
        from_user.amount -= amount;
        to_user.amount += amount;
    }
};

const errors = struct {
    UserExists: error{ UserExists },
    UserNotFound: error{ UserNotFound },
    InsufficientFunds: error{ InsufficientFunds },
    };

pub fn main() void {
    const allocator = std.heap.page_allocator;
    const bank = Bank.init(&allocator) catch {
        std.debug.print("Failed to initialize bank\n", .{});
        return;
    };

    const email1 = "user1@example.com";
    const email2 = "user2@example.com";
    const password = "password";

    bank.addUser(email1, password) catch {
        std.debug.print("Failed to add user 1\n", .{});
        return;
    };
    bank.addUser(email2, password) catch {
        std.debug.print("Failed to add user 2\n", .{});
        return;
    };

    bank.deposit(email1, 1000.0) catch {
        std.debug.print("Failed to deposit to user 1\n", .{});
        return;
    };

    bank.withdraw(email1, 500.0) catch {
        std.debug.print("Failed to withdraw from user 1\n", .{});
        return;
    };

    bank.transfer(email1, email2, 200.0) catch {
        std.debug.print("Failed to transfer from user 1 to user 2\n", .{});
        return;
    };

    const user1 = bank.users.get(email1) orelse null;
    const user2 = bank.users.get(email2) orelse null;

    if (user1 != null) {
        std.debug.print("User 1 balance: {}\n", .{user1.amount});
    }

    if (user2 != null) {
        std.debug.print("User 2 balance: {}\n", .{user2.amount});
    }
}
