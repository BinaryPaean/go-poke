go-poke
=======

A simple, self-hosted tool to check for uptime and response speed of remote websites. Enables black-box monitoring from any host that you can deploy a go program to.

Version
-------
0.0.2

Example Usage:
--------------

    $gopoke http://www.google.com
    {
      "Timestamp": "2014-12-09T15:45:38.370062494-08:00",
      "Host": "BinaryPaean.local",
      "GoPokeVersion": "0.0.2",
      "Results": [
        {
          "Name": "DNS Lookup",
          "Target": "www.google.com",
          "Metrics": {
            "duration": "41.196152ms",
            "hash": "G4BYPCidvH2ineqeH/uGfaeckjpilQgYLMcGrQ435lI="
          },
          "Err": null
        },
        {
          "Name": "HTTP GET",
          "Target": "http://www.google.com",
          "Metrics": {
            "duration": "154.674217ms",
            "hash": "ir6e6SQCNG7JR+EiyTN5lP4zKmFb69iaFS7tWVGTX78="
          },
          "Err": null
        }
      ]
    }

Context:
--------
The idea is to use cron (or similar) to schedule pokes to sites you want to monitor. Gopoke returns metrics such as the speed of the response, and a hash of the contents to allow for basic uptime, response, and content monitoring.

Redirect the output of gopoke to a log, a dashboard, or alert system that understands JSON.

Arguments & Flags:
------------------
(All flags TBD, see below).

 * -c <filepath> or --config <filepath>: read the configuration file at <filepath>
 * -v or --verbose: Output extensive log and debugging information including the chain of response metrics in order of execution.
 * -V or --version: Output gopoke version, then quit.
 * -? or -h or --help: Output basic usage and flag information, then quit.
 * -t or --timeout: Set timeout duration for all actions. If the timeout is exceeded during an action, the "Err" flag for that action is set, and the request is considered to have failed.

Configuration File Format:
--------------------------
TBD. JSON to match output?

Usage:
------
gopoke <flags> [<url1>, ...<urlN>]

Is it Any Good?
---------------
Not really. Right now gopoke has the actions DNS lookup, and HTTP get request. It has the metrics of duration and SHA256 hash. Currently gopoke will always run all actions on the given list of URLs, and all metrics on each action.

Planned Features & TODO
----------------
A later version of gopoke will include a built-in dashboard for better out-of-the-box experience and utility.

Output will eventually be configurable into at least JSON, CSV, and syslog formats.

Flags not yet implemented, will allow user to select subset of actions and metrics.

TODO - How do people plug in new metrics?
TODO - Current action/metric interface is probably too narrow.

No config file handling yet.