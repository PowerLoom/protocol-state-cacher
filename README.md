# Libp2p-submission-sequencer-listener Deployment
Scripts to deploy Sequencer-listener

## Requirements

1. Latest version of `docker` (`>= 20.10.21`) and `docker-compose` (`>= v2.13.0`)
2. At least 4 core CPU, 8GB RAM and 50GB SSD - make sure to choose the correct spec when deploying to Github Codespaces.

## Running the Sequencer Node

Clone the repository against the testnet branch.

`git clone https://github.com/PowerLoom/libp2p-submission-sequencer-listener.git --single-branch powerloom_sequencer_listener && cd powerloom_sequencer_listener`


### Deployment steps

1. Copy `env.example` to `.env`.
    - Ensure the following required variables are filled:
        - `RENDEZVOUS_POINT`: The identifier for locating all relayer peers which are the only way to access the sequencer and submit snapshots.
        - `PROTOCOL_STATE_CONTRACT`: The contract address for the protocol state.
        - `PROST_RPC_URL`: The URL for the PROST RPC service.
        - `DATA_MARKET_ADDRESS`: The contract address of data market this listener is for.

    - Optionally, you may also set the following variables:
        - `REDIS_HOST` & `REDIS_PORT`: The redis server connection url (if you wish to use a separate one).
        - `SLACK_REPORTING_URL`: The reporting url for sending alert notifications.

2. Build the image

   `./build-docker.sh`

3. Run the following command (ideally in a `screen`) and follow instructions

   `./run.sh`

## Troubleshooting
### To be added
### Stopping and Resetting
1. To shutdown services, just press `Ctrl+C` (and again to force).# protocol-state-cacher
