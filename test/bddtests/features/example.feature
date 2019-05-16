#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@example
Feature:

  @example_s1
  Scenario: example
    Given the channel "mychannel" is created and all peers have joined
    And collection config "coll3" is defined for collection "collection3" as policy="OR('Org1MSP.member','Org2MSP.member')", requiredPeerCount=1, maxPeerCount=2, and blocksToLive=3
    And "test" chaincode "examplecc" is installed from path "github.com/trustbloc/example_cc" to all peers
    And "test" chaincode "examplecc" is instantiated from path "github.com/trustbloc/example_cc" on the "mychannel" channel with args "" with endorsement policy "AND('Org1MSP.member','Org2MSP.member')" with collection policy "coll3"
    And chaincode "examplecc" is warmed up on all peers on the "mychannel" channel


    # Test transactions
    When client invokes chaincode "examplecc" with args "del,k1" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "get,k1" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "examplecc" to client equal value ""

    When client invokes chaincode "examplecc" with args "put,k1,20" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "get,k1" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "examplecc" to client equal value "20"

    When client invokes chaincode "examplecc" with args "put,k1,20" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "get,k1" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "examplecc" to client equal value "20-20"

    When client invokes chaincode "examplecc" with args "del,k1" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "get,k1" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "examplecc" to client equal value ""


    # Test private data collections transactions
    When client invokes chaincode "examplecc" with args "putprivate,collection3,pvtKey,pvtVal" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "getprivate,collection3,pvtKey" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "ol_examplecc" to client equal value "pvtVal"

    When client invokes chaincode "examplecc" with args "delprivate,collection3,pvtKey" on the "mychannel" channel
    And we wait 2 seconds
    And client queries chaincode "examplecc" with args "getprivate,collection3,pvtKey" on a single peer in the "peerorg1" org on the "mychannel" channel
    Then response from "ol_examplecc" to client equal value ""