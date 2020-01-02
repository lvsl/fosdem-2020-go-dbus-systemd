# Materials for FOSDEM'20 talk - [Uplift your Linux systems programming skills with systemd and D-Bus](https://fosdem.org/2020/schedule/event/golinux/)

## Agenda:
- [ ] Intro *(5 min)*
  - [ ] About Me ...
      * My 3rd FOSDEM, got impressed a lot by the community!
      * Greatful to org. commette for accepting my presentation
    - [ ] My background with Linux
    - [ ] Why I decided to come here and present at FOSDEM
    - [ ] How I managed to ignore D-Bus and systemd for 10 years?!
      * Also didn't have a TV for 10 years, and bought one recently. It's great!
  - [ ] About This Talk ...
    - [ ] Scope: Systems Programming
      * Def.: "Software that provides services for other (application) software" [[wikipedia](https://en.wikipedia.org/wiki/Systems_programming)]

- [ ] D-Bus *(10 min)*
  - [ ] What is D-Bus?
  - [ ] Contrast and compare it with REST and gRPC
  - [ ] What projects using it (Chrome OS, Firefox, Systemd)
  - [ ] The Good Parts ...
    * Passing File Descriptors as a Native Concept
    * Local Service Discovery
  - [ ] The Ugly Parts ...
    * Object Model 
    * Variable types
    * Discovering API
    * D-Bus daemon quirks
  - [ ] Demo Time
    * Opening a New Tab in Firefox via D-Bus
    * Snooping D-Bus traffic

- [ ] Systemd *(10 min)*
  - [ ] What is Systemd?
  - [ ] How it looks like and why that many features?
    * Every User has a systemd
  - [ ] Demo Time
    * Long-running Background Jobs that Survice Crashes
    * Graceful Termination and Startup
    * Watchdogs and Notifications

- [ ] Q & A *(5 min)*
