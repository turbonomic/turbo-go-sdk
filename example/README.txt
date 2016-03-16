

This is a example of how to write a probe using VMTurbo Go SDK. This example sets up a websocket 
connection as a simulated VMT target to a VMTServer, once the connection is established, a simulated 
Target Supply Chain definition is created and sent to the server. You can refer to the example and 
create any type of supply chain based on your infrastructure.
This default simulated Supply Chain definition contains a seller entity of type Physical Machine
and a buyer entity of type Virtual Machine. 
The default topology which will be discovered by the VMTServer is 2 Physical Machine sellers;
the first PM sells to two different virtual machines and the second PM sells to one virtual machine.



Before running this program:

-Get your local IP address and update the local_IP variable in the main() function
-Get the IP address of the VMTServer host you want to connect to
and update the VMTServer_IP variable in the main() function.
-Create an arbitrary identifier for this simulated target and update the TargetIdentifier variable

To Modify the Default Supply Chain:
-The function createSupplyChain() includes comments explaining how to add entity types  with their 
 respective relationship to other entity types.
 Other available EntityDTO_EntityType which can be added to the supply chain can be found in
 vmturbo-go-sdk/sdk/CommonDTO.pb.go
-The steps are: create a supplyChainNodeBuilder instance with sdk.NewSupplyChainNodeBuilder() , 
 set the supplyChainNodeBuilder object's EntityDTO_EntityType and other properties, add the 
 supplyChainNodeBuilder instance to the supplyChain Builder.

To Modify the Default Topology:
-The method SampleProbe() defines the targets topology (entities and their properties with relationships
 to each other).
-To add entities create new Entity_Params objects with the corresponding fields set and add the entities
 to the entities array. Only one entities array is used in this method.
