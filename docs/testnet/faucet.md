<!--
order: 2
-->

# Faucet

Check how to obtain testnet tokens from the Echelon faucet website {synopsis}

The Echelon Testnet Faucet distributes small amounts of {{ $themeConfig.project.testnet_denom }} to anyone who can provide a valid testnet address for free. Request funds from the faucet either by using the [Keplr Wallet](../guides/keys-wallets/keplr.md) or follow the instructions on this page.

::: tip
Follow the [Metamask](./../guides/keys-wallets/metamask.md), [Keplr](../guides/keys-wallets/keplr.md) or [Keyring](./../guides/keys-wallets/keyring.md) guides for more info on how to setup your wallet account.
:::

## Request tokens

You can request tokens for the testnet by using the Echelon [faucet](https://faucet.echelon.dev).
Simply fill in your address on the input field in Bech32 (`echelon1...`) or Hex (`0x...`) format.

::: warning
If you use your Bech32 address, make sure you input the [account address](./../basics/accounts.md#addresses-and-public-keys) (`echelon1...`) and **NOT** the validator operator address (`echelonvaloper1...`)
:::

![faucet site](./img/faucet_web_page.png)

## Rate limits

To prevent the faucet account from draining the available funds, the Echelon testnet faucet
imposes a maximum number of request for a period of time. By default the faucet service accepts 1
request per day per address. All addresses **must** be authenticated using
ReCAPTCHA before requesting tokens.

## Amount

For each request, the faucet transfers 1 {{ $themeConfig.project.testnet_denom }} to the given address.

## Faucet Addresses

The public faucet addresses for the testnet are:

- **Hex**: [`0xBaE9A7A2210F94511F5050348251d0d7113E2cE3`](https://evm.echelon.dev/address/0xBaE9A7A2210F94511F5050348251d0d7113E2cE3/transactions)
- **Bech32**: [`echelon1ht560g3pp729z86s2q6gy5ws6ugnut8r4uhyth`](https://testnet.mintscan.io/echelon/account/echelon1ht560g3pp729z86s2q6gy5ws6ugnut8r4uhyth)
