# Matter Server UI Home Assistant Add-on

## Overview

The matter-server-ui add-on for Home Assistant provides a user-friendly, web-based interface to interact with python-matter-server. It offers a convenient way to monitor and debug Matter devices within your Home Assistant environment. Utilizing Home Assistant's ingress feature, it seamlessly integrates into your Home Assistant dashboard, allowing for easy access and real-time insights.

## Features

- **Real-Time Monitoring**: View live updates and statuses of Matter nodes, including attributes, events, and network status.
- **Interactive Control**: Send commands and receive responses from Matter devices directly through the UI.
- **Integrated Debugging**: Access detailed logs and debugging tools for in-depth network analysis.
- **Ingress Integration**: Fully integrated into Home Assistant's UI, providing easy and secure access without additional port configuration.

## Usage

### Accessing the UI:

Once the add-on is installed, it will appear in the Home Assistant dashboard.
Click on the matter-server-ui add-on to open the interface.

### Configuration:

All configurations can be done through the add-on's UI.

Configuring matter-server-ui in Home Assistant requires special attention, especially when dealing with WebSocket connections and TLS (Transport Layer Security). Below are detailed instructions on how to set the scheme, host, port, and path in the matter-server-ui settings to ensure a secure and functioning connection.

#### Understanding the Settings Fields

- **Scheme**: This can be either ws (WebSocket) or wss (WebSocket Secure).
- **Host**: The network address of your python-matter-server.
- **Port**: The network port on which your python-matter-server is listening.
- **Path**: The path used by the WebSocket to connect to the server.

#### Configuration Scenarios

##### Without TLS

- **Scheme & Path**: If TLS is not enabled in your Home Assistant setup, you can usually leave the scheme as ws and the path as default.
- **Host & Port**: Enter the host and port of your python-matter-server.

##### With TLS

When TLS is enabled, direct connections to a non-TLS WebSocket are blocked due to browser security restrictions. In this case, you need to adjust the settings as follows:

- **Scheme**: Change the scheme to wss to indicate a secure WebSocket connection.
- **Host**: Your home assistant host
- **Port**: Your home assistant port (likely 443)
- **Path**: Set the path to <homeassistant add-on ingress path>/wsproxy?host=<targethost>&port=<targetport>. The default port for WebSocket is usually 5580 if not specified.
  - To find the Home Assistant add-on ingress path, you may need to inspect network requests using your browser's developer tools while interacting with the matter-server-ui add-on. Look for requests sent by this add-on to determine the ingress path.

#### Example Configuration

If your Home Assistant is running with TLS and your python-matter-server is on the same network host at port 5580, your configuration might look like this:

- **Scheme**: wss
- **Host**: home-assistant.local
- **Port**: 443
- **Path**: /api/hassio_ingress/abc123/wsproxy?host=my-matter-server.local
  - Replace my-matter-server.local with the actual host of your python-matter-server and /api/hassio_ingress/abc123 with the actual ingress path of your matter-server-ui add-on


### Navigation:

- **Server Info**: Check the current status and details of the connected python-matter-server.
- **Nodes**: View and interact with Matter nodes connected to your network.
- **Event Log**: Monitor real-time events and activities from your Matter devices.
- **Settings**: Adjust and save your connection settings to the python-matter-server.
