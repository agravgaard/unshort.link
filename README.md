# Unshort.link

*This is a fork of [unshort.link](https://github.com/simonfrey/unshort.link).*
*The upstream was archived. This fork attempts to keep the project and it's dependencies up-to-date.*

Prevent short link services from tracking you by un shortening the urls for your. Try it on [unshort.link](https://unshort.link)

## Features

- Access short links for you to prevent the short link providers to track you
- Check links against a blacklist to prevent access to a harmful website hidden behind a short link
- Remove known tracking parameters from the urls behind the short links (e.g. the facebook tracking parameter `utm_source`)
- Remove as many url parameters as possible by keeping the same website result. This helps to remove tracking parameters that are so far unknown

## Contributors

Thanks to all the following contributors for their work on unshort.link!

- [simonfrey](https://github.com/simonfrey) for [unshort.link](https://github.com/simonfrey/unshort.link) - the project this repo was forked from.
- [shayne](https://github.com/shayne) for the fix in the webextension not loading the providers from the custom server (Mar 2020)
- [roket1428](https://github.com/roket1428) for the logo and the dark design (Jan 2020)
- [cyantarek](https://github.com/cyantarek) for adding the makefile 
- [Jakob-em](https://github.com/Jakob-em) for periodically reloading the blacklist & UI improvements (Jan 2020)
- [billcobbler](https://github.com/billcobbler) for the dockerization of the server (Jan 2020)
- [madstk1](https://github.com/madstk1) for the new bootstrap based frontend design (Dec 2019)
- [dkter](https://github.com/dkter) for the bugfix bugfix of white text on white ground (Oct 2019)
