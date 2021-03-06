god
===

god is a scalable, performant, persistent in memory data structure server. It allows massively distributed applications to update and fetch common data in a structured and sorted format.

Its main inspirations are Redis and Chord/DHash. Like Redis it focuses on performance, ease of use and a small, simple yet powerful feature set, while from the Chord/DHash projects it inherits scalability, redundancy and transparent failover behaviour.

# Try it out

<code>go get github.com/zond/god/god_server</code>, run <code>god_server</code>, browse to <a href="http://localhost:9192/">http://localhost:9192/</a>.

# Documents

HTML documentation: http://zond.github.com/god/

godoc documentation: http://go.pkgdoc.org/github.com/zond/god

# TODO

* Docs
 * Add illustrations to the usage manual
* Benchmark
 * Consecutively start 1-20 ec2 small instances and benchmark against each size
  * Need 20 machines of equal and constant performance. Is anyone willing to lend me this for few days of benchmarking?
 * Add benchmark results to docs
