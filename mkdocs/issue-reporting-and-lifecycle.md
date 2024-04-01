# Support and Issue Reporting

## Issue Creation Guidelines

We welcome issues of all kinds including feature requests, bug reports, and
general questions. Below you'll find checklists with guidelines for well-formed
issues of each type.

We encourage opening new issues rather than commenting on closed issues if a problem has not been completely solved or causes a regression. This ensures we are able to triage it effectively.

### [Bug Reports](https://github.com/bschaatsbergen/terraform-provider-assert/issues/new?template=Bug_Report.md)

- __Test against the latest release__: Make sure you test against the latest
   released version. It is possible we already fixed the bug you're experiencing.

- __Search for possible duplicate reports__: It's helpful to keep bug
   reports consolidated to one thread, so do a quick search on existing bug
   reports to check if anybody else has reported the same thing. You can [scope
      searches by the label "bug"](https://github.com/bschaatsbergen/terraform-provider-assert/issues?q=is%3Aopen+is%3Aissue+label%3Abug) to help narrow things down.

- __Include steps to reproduce__: Provide steps to reproduce the issue,
   along with your `.tf` files, with secrets removed, so we can try to
   reproduce it. Without this, it makes it much harder to fix the issue.

- __For panics, include `crash.log`__: If you experienced a panic, please
   create a [gist](https://gist.github.com) of the *entire* generated crash log
   for us to look at. Double-check check no sensitive items were in the log.

### [Feature Requests](https://github.com/bschaatsbergen/terraform-provider-assert/issues/new?labels=enhancement&template=Feature_Request.md)

- __Search for possible duplicate requests__: It's helpful to keep requests
   consolidated to one thread, so do a quick search on existing requests to
   check if anybody else has reported the same thing. You can [scope searches by
      the label "enhancement"](https://github.com/bschaatsbergen/terraform-provider-assert/issues?q=is%3Aopen+is%3Aissue+label%3Aenhancement) to help narrow things down.

- __Include a use case description__: In addition to describing the
   behavior of the feature you'd like to see added, it's helpful to also lay
   out the reason why the feature would be important and how it would benefit
   Terraform users.
