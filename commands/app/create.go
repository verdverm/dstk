package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func CreateDstkScalaSparkApp(appname string) {
	cwd, err := os.Getwd()
	panicErr(err)

	appdir := fmt.Sprintf("%s/%s", cwd, appname)
	srcdir := fmt.Sprintf("%s/src/%s", appdir, appname)
	appfn := srcdir + "/" + appname + ".scala"

	os.MkdirAll(srcdir, 0755)

	tmpl, err := template.New("appfile").Parse(scala_template)
	panicErr(err)

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tmpl.Execute(w, map[string]string{
		"APPNAME":  appname,
		"APPCLASS": strings.Title(appname),
	})
	panicErr(err)
	w.Flush()

	ioutil.WriteFile(appfn, buf.Bytes(), 0644)
	ioutil.WriteFile(appdir+"/MANIFEST.MF", []byte(manifest_template), 0644)
}

var scala_template = `/*
 * License and Copyright...
 */

package {{.APPNAME}}

import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._
import org.apache.spark.SparkConf


object {{.APPCLASS}} {
  def main(args: Array[String]) {
    val conf = new SparkConf().setAppName("Spark Scala {{.APPCLASS}}")
    val spark = new SparkContext(conf)

    /* the following is a sorted wordcount example */

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
`

var manifest_template = `Manifest-Version: 1.0
Class-Path: scala-library.jar spark-assembly-1.0.2-hadoop2.2.0.jar wordcount.jar
`
