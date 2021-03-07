# pgtypes

This package provides wrapper types for converting various go types to and
from their postgresql representation via the standard `database/sql` interfaces.
The original implementation of these types came from the `github.com/lib/pq` package,
which you can still use to gain access to these types. The reason to use `ethanpailes/pgtypes`
instead is that `lib/pq` is that `lib/pq` is no longer maintained and bundles an
entire postgres driver. `ethanpailes/pgtypes` just provides the conversion types.

# Licensing & Copyright

The copyright for most of this code is held by the `lib/pq` contributors and
Blake Mizerany and licenced under the MIT license. The copyright for some recent
changes is held by Ethan Pailes and `ethanpailes/pgtypes` contributors under the
same license.
