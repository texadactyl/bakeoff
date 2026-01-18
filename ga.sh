#!/usr/bin/env bash

set -euo pipefail
unset JAVA_TOOL_OPTIONS

LOGFILE=ga.log
exec > >(tee "$LOGFILE") 2>&1

timestamp() {
    date -u +"%Y-%m-%dT%H:%M:%SZ"
}

section() {
    echo
    echo "=== $1 ==="
}

kv() {
    printf "%-24s %s\n" "$1:" "$2"
}

timestamp
section "ENVIRONMENT"

kv "OS"        "${RUNNER_OS:-unknown}"
kv "ARCH"      "$(uname -m)"
kv "KERNEL"    "$(uname -sr)"

ARCH=$(uname -m)

# ----------------------------
# Architecture-specific flags
# ----------------------------
# Use GitHub Actions RUNNER_ARCH if available, fallback to uname
ARCH="${RUNNER_ARCH:-$(uname -m)}"

# Normalize to lowercase for consistent matching
ARCH=$(echo "$ARCH" | tr '[:upper:]' '[:lower:]')

case "$ARCH" in
    x64|x86_64|amd64)
        C_CPU="-march=x86-64 -mtune=generic"
        D_CPU="-mcpu=generic"
        ZIG_CPU="-mcpu=x86_64"
        ;;
    arm64|aarch64)
        C_CPU="-march=armv8-a"
        D_CPU="-mcpu=generic"
        ZIG_CPU="-mcpu=generic"
        ;;
    *)
        echo "WARNING: Unknown architecture '$ARCH', using baseline"
        C_CPU=""
        D_CPU=""
        ZIG_CPU="-mcpu=baseline"
        ;;
esac

COMMON_CFLAGS="-O3 -ffast-math"
COMMON_DFLAGS="-O3 -release -ffast-math"
COMMON_ZIGFLAGS="-O ReleaseFast"

# ============================
# C
# ============================
section "C"

kv "Compiler" "$(gcc -dumpmachine)"
kv "Version"  "$(gcc --version | head -n 1)"

gcc $COMMON_CFLAGS $C_CPU main.c -o c.out
./c.out

# ============================
# D (LDC2)
# ============================
section "D (ldc2)"

kv "Compiler" "ldc2"
kv "Version"  "$(ldc2 --version 2>&1 | grep -m1 'LDC' | tr '\n' ' ')"

ldc2 $COMMON_DFLAGS $D_CPU main.d -of=d.out
./d.out

# ============================
# Zig
# ============================
section "Zig"

kv "Compiler" "zig"
kv "Version"  "$(zig version)"

zig build-exe main.zig \
    $COMMON_ZIGFLAGS \
    $ZIG_CPU \
    -femit-bin=zig.out

./zig.out

# ============================
# Go
# ============================
section "Go"

kv "Compiler" "go"
kv "Version"  "$(go version)"

go build -trimpath -ldflags="-s -w" -o go.out main.go
./go.out

# ============================
# Java
# ============================
section "Java Compile"

kv "javac" "$(javac --version)"
javac -Xlint -Werror main.java

section "HotSpot -server"
kv "java" "$(java --version | head -n 1)"
java -server main

section "HotSpot -server +TieredCompilation"
java -server -XX:+TieredCompilation main

section "DONE"
timestamp

