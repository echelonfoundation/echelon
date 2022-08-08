<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog
Echelon v1.2.0 - WIP


Echelon v1.1.5

- Updated protobuf building
- Added x/vrf API and gRPC endpoints
- Preparation update for v2.0.0
- Code cleanup

Echelon v1.1.4

- Upgraded Cosmos SDK to latest v0.45.6 *See Cosmos-SDK changelog for breaking changes*
- Upgraded Ethermint v0.14.1
- Upgraded Tendermint v0.34.19
- Added spendable balances API endpoint (/cosmos/bank/v1beta1/spendable_balance/)
- Fixed protobuf duplicate warnings
- Updated initial peers list configuration
- Included genesis.json in main repo
- x/erc20 module upgraded to use better regex
- Updated to use Go Lang v1.18+