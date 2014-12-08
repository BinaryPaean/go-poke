go-poke
=======

A simple, self-hosted tool to check for uptime and response speed of remote websites.

Version
-------
0.0.2

Example Usage:
--------------

    ./gopoke http://www.google.com
2014/12/08 10:33:26 {
  "Timestamp": "2014-12-08T10:33:26.556042835-08:00",
  "Host": "BinaryPaean.local",
  "Version": "0.0.2",
  "Results": [
    {
      "Name": "DNS Lookup",
      "Target": "www.google.com",
      "Metrics": {
        "duration": "47.930684ms",
        "hash": "OeJMhuIIuxlHldMAxoYKuV8gMHakELCvN7sWumbYe3c="
      },
      "Err": null
    },
    {
      "Name": "HTTP GET",
      "Target": "http://www.google.com",
      "Metrics": {
        "duration": "348.449618ms",
        "hash": "OigC9mo/TCIv0KYsPxNmQBKQvdYZhmK2OuPLaW80FiQ="
      },
      "Err": null
    }
  ]
}

Context:
--------
The idea is you use cron (or similar) to schedule pokes to sites you want to monitor. Gopoke returns metrics such as the speed of the response, and a hash of the contents to allow for basic uptime, response, and content monitoring.

Redirect the output of gopoke to a log, a dashboard, or any system that understands json for use and abuse as you see fit.

 A later version of gopoke will include a built-in dashboard for better out-of-the-box experience and utility.

Arguments & Flags:
------------------
 * -c <filepath> or --config <filepath>: read the configuration file at <filepath>
 * -v or --verbose: Output extensive log and debugging information including the chain of response metrics in order of execution.

Configuration File Format:
--------------------------
JSON to match output?

TBD - specify set of metrics to run, or always run the full set?
TBD - How do people plug in new metrics? "Stupidest thing that could work":
Users just add metrics in source, program always runs all metrics on all targets.

Usage:
------
gopoke <flags> [<url1>, ...]

Is it Any Good?
---------------
Barely. Right now gopoke has the actions DNS lookup, and HTTP get request. It has the metrics of duration and SHA256 hash. For a list of URLs it will always run all actions, and all metrics on each action until command line flags are implemented.