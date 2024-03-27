/*
 * Copyright (c) 2014, Oracle America, Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *
 *  * Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 *  * Neither the name of Oracle nor the names of its contributors may be used
 *    to endorse or promote products derived from this software without
 *    specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
 * THE POSSIBILITY OF SUCH DAMAGE.
 */

package org.sample;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.ThreadLocalRandom;
import java.util.concurrent.TimeUnit;

import org.apache.commons.lang3.RandomStringUtils;
import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Level;
import org.openjdk.jmh.annotations.Measurement;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.annotations.Param;
import org.openjdk.jmh.annotations.Scope;
import org.openjdk.jmh.annotations.Setup;
import org.openjdk.jmh.annotations.State;
import org.openjdk.jmh.annotations.TearDown;
import org.openjdk.jmh.annotations.Warmup;
import org.openjdk.jmh.infra.Blackhole;

public class MyBenchmark {

    // @Benchmark
    // @BenchmarkMode(Mode.AverageTime)
    // @Warmup(iterations = 10, time = 500, timeUnit = TimeUnit.MILLISECONDS)
    // @Measurement(iterations = 30, time = 1000, timeUnit = TimeUnit.MILLISECONDS)
    // @OutputTimeUnit(TimeUnit.MILLISECONDS)
    // public void testWithSingleThread() {
    // CacheImpl cache = new CacheImpl();
    // for (int i = 0; i < 100000; i++) {
    // var k = String.valueOf(ThreadLocalRandom.current().nextLong() % 100);
    // var v = String.valueOf(ThreadLocalRandom.current().nextLong() % 100);
    // cache.put(k, v);
    // }
    // }

    @State(Scope.Thread)
    public static class Plan {
        @Param({ "1", "2", "4", "10" })
        private int threadPoolSize;

        private String[] strings;

        @Setup(Level.Iteration)
        public void setup() {
            var nStrings = 100;
            strings = new String[nStrings];
            for (int i = 0; i < nStrings; i++) {
                strings[i] = RandomStringUtils.randomAlphabetic(ThreadLocalRandom.current().nextInt(3, 10));
            }
        }
    }

    @Benchmark
    @BenchmarkMode(Mode.AverageTime)
    @Warmup(iterations = 3, time = 500, timeUnit = TimeUnit.MILLISECONDS)
    @Measurement(iterations = 5, time = 1000, timeUnit = TimeUnit.MILLISECONDS)
    @OutputTimeUnit(TimeUnit.MILLISECONDS)
    public void testWithMultipleThread(Plan plan, Blackhole bh) throws InterruptedException, ExecutionException {
        ExecutorService executorService = Executors.newFixedThreadPool(plan.threadPoolSize);
        try {
            CacheImpl cache = new CacheImpl();
            List<Callable<Boolean>> tasks = new ArrayList<>();
            for (int i = 0; i < 1000000; i++) {
                var r = ThreadLocalRandom.current().nextInt(100);
                if (r < 10) {
                    tasks.add(() -> {
                        var k = plan.strings[ThreadLocalRandom.current().nextInt(100)];
                        var v = plan.strings[ThreadLocalRandom.current().nextInt(100)];
                        cache.put(k, v);
                        return true;
                    });
                } else {
                    tasks.add(() -> {
                        var k = plan.strings[ThreadLocalRandom.current().nextInt(100)];
                        var v = cache.get(k);
                        bh.consume(v);
                        return true;
                    });
                }
            }
            List<Future<Boolean>> futures = executorService.invokeAll(tasks);
            for (var f : futures) {
                f.get();
            }
        } finally {
            executorService.shutdown();
        }
    }
}
