/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wordcount

import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._
import org.apache.spark.SparkConf


object WordCount {
  def main(args: Array[String]) {
    val conf = new SparkConf().setAppName("Spark WordCount")
    val spark = new SparkContext(conf)

    val file = spark.textFile(args(0))

    val counts = file.flatMap(line => line.split(" "))
                     .map(word => (word, 1))
                     .reduceByKey((a, b) => a + b)

    val swapped = counts.map(item => item.swap)
    val sorted = swapped.sortByKey(false) // false == descening order
    val top = sorted.take(100)

    val results = spark.parallelize(top)
    results.saveAsTextFile(args(1))

    spark.stop()
  }
}
