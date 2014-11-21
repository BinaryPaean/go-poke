go-poke
=======

A simple, self-hosted tool to check for uptime and response speed of remote websites.
v0.0.1

Example Usage:
==============

    $gopoke http://www.google.com
    ${ "Poke":
      { "Timestamp": "Fri Nov 21 11:13:31 PST 2014",
        "Target": "http://www.google.com",
        "Host": "pokey.mybox.com",
        "Version": "0.0.1"
      },
      "Latency":
      { "DNS" : 30,
        "Time to First Byte" : 90,
        "Time to Completion" : 130,
      },
      "Response":
        { "Size" : 870,
          "Code" : 200,
          "Headers" : ...,
          "SHA-2/224" : d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f
        }
    }
    $

Context:
========
The idea is you use crontab to schedule pokes to sites you want to monitor.
Redirect the output of gopoke to a file, a dashboard,
or any system that understands json. A later version of gopoke will include a
built-in dashboard for better out-of-the-box experience and utility.

Arguments & Flags:
==================
 * -c <filepath> or --config <filepath>: read the configuration file at <filepath>
 * -v or --verbose: Output extensive log and debugging information including the
   chain of response metrics in order of execution.

Configuration File Format:
==========================
JSON to match output?

TBD - specify set of metrics to run, or always run the full set?
TBD - How do people plug in new metrics? "Stupidest thing that could work":
Users just add metrics in source, program always runs all metrics on all targets.

Usage:
======


Is it Any Good?
===============
Not yet, no.  
In fact due to [readme-driven-development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html)
the rest of this project doesn't exist yet.
