## 2.0.0

This is a breaking-change release that updates ngrok-go to a new, simplified
API.

Enhancements:
- Dramatically simplified the API to remove many overlapping options, options
  that are now deprecated, unnecessary convenience functions, and more.
- Simplified the API by removing all protocol-specific behaviors (which have
  all been moved to Traffic Policy).
- Removed the config package. All of its options are now folded into the
  top-level package or removed because they were migrated into Traffic Policy.
- Updates the API to use new ngrok terminology of Agents, Endpoints Upstreams,
  and Traffic Policy.
- Removes functionality that is now deprecated (like labeled tunnels).
- Added support for agent-based TLS termination and Mutual TLS termination.
- Added support for full TLS control over forwarding to the upstream.
- Added support for full control over dialing the upstream.
- Removed a bespoke logging interface in favor of `log/slog`.
- Removed the prototype policy package that was not well supported.
- Separated out a concept of an Agent from its Session which were previously
  co-mingled.
- Added integration tests.

## 1.12.1

Fixes:

- Fix reported version of the ngrok-go sdk.

## 1.12.0

Breaking changes:

- Renames pre-release option `WithAllowsPooling` to `WithPoolingEnabled`

## 1.11.0

Enhancements:

- Adds `WithAllowsPooling` option, allowing the endpoint to pool with other endpoints with the same host/port/binding
- Adds `WithURL` option, specifying a URL for the endpoint
- Adds `WithTrafficPolicy` option, applying Traffic Policy to the endpoint

Changes:

- Deprecates `WithPolicyString` in favor of `WithTrafficPolicy`

## 1.10.0

Enhancements:

- Adds fasthttp example
- Adds `WithBindings` option
- Adds support for `TrafficPolicy` field
 
Changes:

- Replace log adapter module license symlinks with full files
- Send `Policy` to the backend as a `TrafficPolicy` string
- unsafe.Pointer -> atomic.Pointer

Bug fixes:

- Propagate half-closes correctly in forward

## 1.9.1

Bug fixes:

- Protect against writing to closed channel on shutdown

## 1.9.0
Enhancements:

- Adds `WithAdditionalServers` and `WithMultiLeg` options

## 1.8.1
Enhancements:

- Provides access to structs for building a Traffic Policy configuration

Breaking changes:

- Renames pre-release option `WithPolicyConfig` to `WithPolicyString`
- Changes signature of `WithPolicy` option to accept the newly exposed structs for building a Traffic Policy configuration

## 1.8.0
- Adds the `WithPolicy` and `WithPolicyConfig` options for applying a Traffic Policy to an endpoint.

## 1.7.0

- Adds the `WithAppProtocol` option for labeled listeners and HTTP endpoints.

  This provides a protocol hint that can be used to enable support for HTTP/2 to
  the backend service.

## 1.6.0

- Adds support for remote stop of listener.

## 1.5.1

- Adds TLS Renegotiation to the backend `tls.Config`.

## 1.5.0

- Added new forwarding API. See `[Session].ListenAndForward` and `[Session].ListenAndServeHTTP`.
- Deprecates `WithHTTPServer` and `WithHTTPHandler`. Use `[Session].ListenAndServeHTTP` instead.

## 1.4.0

- Switch to `connect.ngrok-agent.com:443` as the default server address
- Add nicer error types that expose the ngrok error code

## 1.0.0 (2023-01-10)

Enhancements:

- Initial release
