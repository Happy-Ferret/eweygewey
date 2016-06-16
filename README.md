EweyGewey
=========

EweyGewey is an OpenGL GUI library written in the [Go][golang] programming language
that is designed to be integrated easily into an OpenGL application.

UNDER CONSTRUCTION
==================

At present, it is very much in an alpha stage with new development adding in
features, widgets and possibly API breaks.

Screenshots
-----------

Here's some of what's available right now in the [basic example][basic_example]:

![basic_ss][basic_ss]


Requirements
------------

* [Mathgl][mgl] - for 3d math
* [Freetype][ftgo] - for dynamic font texture generation
* [Fizzle][fizzle] - provides an OpenGL 3/es2/es3 abstraction
* [GLFW][glfw-go] (v3.1) - currently GLFW is the only 'host' support for input

Additionally, a backend graphics provider needs to be used. At present, fizzle
supports the following:

* [Go GL][go-gl] - pre-generated OpenGL bindings using their glow project
* [opengles2][go-gles] - Go bindings to the OpenGL ES 2.0 library

These are included when the `graphicsprovider` subpackage is used and direct
importing is not required.

Installation
------------

The dependency Go libraries can be installed with the following commands.

```bash
go get github.com/go-gl/glfw/v3.1/glfw
go get github.com/go-gl/mathgl/mgl32
go get github.com/golang/freetype
go get github.com/tbogdala/fizzle
```

An OpenGL library will also be required for desktop applications; install
the OpenGL 3.3 library with the following command:

```bash
go get github.com/go-gl/gl/v3.3-core/gl
```

If you're compiling for Android/iOS, then you will need an OpenGL ES library,
and that can be installed with the following command instead:

```bash
go get github.com/remogatto/opengles2
```

This does assume that you have the native GLFW 3.1 library installed already
accessible to Go tools. This should be the only native library needed.

Current Features
----------------

* Basic windowing system
* Basic theming support
* Basic input support that detects mouse clicks and double-clicks
* Widgets:
    * Text
    * Buttons

examples
--------

Currently there is only one example, but all of the examples will likely use
the `exampleapp.go`, "engine-in-a-file" support code found in the example directory.
Using this just provides some quick scaffolding

TODO
----

The following need to be addressed in order to start releases:

* widgets use a lot of pixel sizing where windows use relative sizing; this
  creates a problem where widgets don't resize to the new window space appropriately.
* more widgets:
    * single/multi-line text editbox
    * combobox
    * images
    * custom widgets
    * scroll bars
* documentation
* samples


LICENSE
=======

EweyGewey is released under the BSD license. See the [LICENSE][license-link] file for more details.

Fonts in the `examples/assets` directory are licensed under the [SIL OFL][sil_ofl] open font license.

[golang]: https://golang.org/
[fizzle]: https://github.com/tbogdala/fizzle
[glfw-go]: https://github.com/go-gl/glfw
[mgl]: https://github.com/go-gl/mathgl
[ftgo]: https://github.com/golang/freetype
[imgui]: https://github.com/ocornut/imgui
[sil_ofl]: http://scripts.sil.org/cms/scripts/page.php?site_id=nrsi&id=OFL
[license-link]: https://raw.githubusercontent.com/tbogdala/eweygewey/master/LICENSE
[basic_ss]: https://raw.githubusercontent.com/tbogdala/cubez/master/examples/screenshots/basic_ss_0.jpg
[basic_example]: https://raw.githubusercontent.com/tbogdala/cubez/master/examples/basic
