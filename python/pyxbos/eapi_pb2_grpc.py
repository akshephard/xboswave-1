# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import eapi_pb2 as eapi__pb2


class WAVEStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateEntity = channel.unary_unary(
        '/mqpb.WAVE/CreateEntity',
        request_serializer=eapi__pb2.CreateEntityParams.SerializeToString,
        response_deserializer=eapi__pb2.CreateEntityResponse.FromString,
        )
    self.CreateAttestation = channel.unary_unary(
        '/mqpb.WAVE/CreateAttestation',
        request_serializer=eapi__pb2.CreateAttestationParams.SerializeToString,
        response_deserializer=eapi__pb2.CreateAttestationResponse.FromString,
        )
    self.PublishEntity = channel.unary_unary(
        '/mqpb.WAVE/PublishEntity',
        request_serializer=eapi__pb2.PublishEntityParams.SerializeToString,
        response_deserializer=eapi__pb2.PublishEntityResponse.FromString,
        )
    self.PublishAttestation = channel.unary_unary(
        '/mqpb.WAVE/PublishAttestation',
        request_serializer=eapi__pb2.PublishAttestationParams.SerializeToString,
        response_deserializer=eapi__pb2.PublishAttestationResponse.FromString,
        )
    self.AddAttestation = channel.unary_unary(
        '/mqpb.WAVE/AddAttestation',
        request_serializer=eapi__pb2.AddAttestationParams.SerializeToString,
        response_deserializer=eapi__pb2.AddAttestationResponse.FromString,
        )
    self.LookupAttestations = channel.unary_unary(
        '/mqpb.WAVE/LookupAttestations',
        request_serializer=eapi__pb2.LookupAttestationsParams.SerializeToString,
        response_deserializer=eapi__pb2.LookupAttestationsResponse.FromString,
        )
    self.ResyncPerspectiveGraph = channel.unary_unary(
        '/mqpb.WAVE/ResyncPerspectiveGraph',
        request_serializer=eapi__pb2.ResyncPerspectiveGraphParams.SerializeToString,
        response_deserializer=eapi__pb2.ResyncPerspectiveGraphResponse.FromString,
        )
    self.SyncStatus = channel.unary_unary(
        '/mqpb.WAVE/SyncStatus',
        request_serializer=eapi__pb2.SyncParams.SerializeToString,
        response_deserializer=eapi__pb2.SyncResponse.FromString,
        )
    self.WaitForSyncComplete = channel.unary_stream(
        '/mqpb.WAVE/WaitForSyncComplete',
        request_serializer=eapi__pb2.SyncParams.SerializeToString,
        response_deserializer=eapi__pb2.SyncResponse.FromString,
        )
    self.BuildRTreeProof = channel.unary_unary(
        '/mqpb.WAVE/BuildRTreeProof',
        request_serializer=eapi__pb2.BuildRTreeProofParams.SerializeToString,
        response_deserializer=eapi__pb2.BuildRTreeProofResponse.FromString,
        )
    self.VerifyProof = channel.unary_unary(
        '/mqpb.WAVE/VerifyProof',
        request_serializer=eapi__pb2.VerifyProofParams.SerializeToString,
        response_deserializer=eapi__pb2.VerifyProofResponse.FromString,
        )
    self.ListLocations = channel.unary_unary(
        '/mqpb.WAVE/ListLocations',
        request_serializer=eapi__pb2.ListLocationsParams.SerializeToString,
        response_deserializer=eapi__pb2.ListLocationsResponse.FromString,
        )
    self.Inspect = channel.unary_unary(
        '/mqpb.WAVE/Inspect',
        request_serializer=eapi__pb2.InspectParams.SerializeToString,
        response_deserializer=eapi__pb2.InspectResponse.FromString,
        )
    self.ResolveHash = channel.unary_unary(
        '/mqpb.WAVE/ResolveHash',
        request_serializer=eapi__pb2.ResolveHashParams.SerializeToString,
        response_deserializer=eapi__pb2.ResolveHashResponse.FromString,
        )
    self.EncryptMessage = channel.unary_unary(
        '/mqpb.WAVE/EncryptMessage',
        request_serializer=eapi__pb2.EncryptMessageParams.SerializeToString,
        response_deserializer=eapi__pb2.EncryptMessageResponse.FromString,
        )
    self.DecryptMessage = channel.unary_unary(
        '/mqpb.WAVE/DecryptMessage',
        request_serializer=eapi__pb2.DecryptMessageParams.SerializeToString,
        response_deserializer=eapi__pb2.DecryptMessageResponse.FromString,
        )
    self.CreateNameDeclaration = channel.unary_unary(
        '/mqpb.WAVE/CreateNameDeclaration',
        request_serializer=eapi__pb2.CreateNameDeclarationParams.SerializeToString,
        response_deserializer=eapi__pb2.CreateNameDeclarationResponse.FromString,
        )
    self.ResolveName = channel.unary_unary(
        '/mqpb.WAVE/ResolveName',
        request_serializer=eapi__pb2.ResolveNameParams.SerializeToString,
        response_deserializer=eapi__pb2.ResolveNameResponse.FromString,
        )
    self.MarkEntityInteresting = channel.unary_unary(
        '/mqpb.WAVE/MarkEntityInteresting',
        request_serializer=eapi__pb2.MarkEntityInterestingParams.SerializeToString,
        response_deserializer=eapi__pb2.MarkEntityInterestingResponse.FromString,
        )
    self.ResolveReverseName = channel.unary_unary(
        '/mqpb.WAVE/ResolveReverseName',
        request_serializer=eapi__pb2.ResolveReverseNameParams.SerializeToString,
        response_deserializer=eapi__pb2.ResolveReverseNameResponse.FromString,
        )
    self.Revoke = channel.unary_unary(
        '/mqpb.WAVE/Revoke',
        request_serializer=eapi__pb2.RevokeParams.SerializeToString,
        response_deserializer=eapi__pb2.RevokeResponse.FromString,
        )
    self.CompactProof = channel.unary_unary(
        '/mqpb.WAVE/CompactProof',
        request_serializer=eapi__pb2.CompactProofParams.SerializeToString,
        response_deserializer=eapi__pb2.CompactProofResponse.FromString,
        )


class WAVEServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateEntity(self, request, context):
    """Create a new WAVE entity, but do not publish it
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateAttestation(self, request, context):
    """Create a WAVE attestation, both the source and destination entities must
    be published
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PublishEntity(self, request, context):
    """Publish the given entity
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PublishAttestation(self, request, context):
    """Publish an attestation
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AddAttestation(self, request, context):
    """Add an attestation to the given perspective graph
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LookupAttestations(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResyncPerspectiveGraph(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SyncStatus(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def WaitForSyncComplete(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def BuildRTreeProof(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def VerifyProof(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListLocations(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Inspect(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResolveHash(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def EncryptMessage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DecryptMessage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNameDeclaration(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResolveName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def MarkEntityInteresting(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResolveReverseName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Revoke(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CompactProof(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_WAVEServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateEntity': grpc.unary_unary_rpc_method_handler(
          servicer.CreateEntity,
          request_deserializer=eapi__pb2.CreateEntityParams.FromString,
          response_serializer=eapi__pb2.CreateEntityResponse.SerializeToString,
      ),
      'CreateAttestation': grpc.unary_unary_rpc_method_handler(
          servicer.CreateAttestation,
          request_deserializer=eapi__pb2.CreateAttestationParams.FromString,
          response_serializer=eapi__pb2.CreateAttestationResponse.SerializeToString,
      ),
      'PublishEntity': grpc.unary_unary_rpc_method_handler(
          servicer.PublishEntity,
          request_deserializer=eapi__pb2.PublishEntityParams.FromString,
          response_serializer=eapi__pb2.PublishEntityResponse.SerializeToString,
      ),
      'PublishAttestation': grpc.unary_unary_rpc_method_handler(
          servicer.PublishAttestation,
          request_deserializer=eapi__pb2.PublishAttestationParams.FromString,
          response_serializer=eapi__pb2.PublishAttestationResponse.SerializeToString,
      ),
      'AddAttestation': grpc.unary_unary_rpc_method_handler(
          servicer.AddAttestation,
          request_deserializer=eapi__pb2.AddAttestationParams.FromString,
          response_serializer=eapi__pb2.AddAttestationResponse.SerializeToString,
      ),
      'LookupAttestations': grpc.unary_unary_rpc_method_handler(
          servicer.LookupAttestations,
          request_deserializer=eapi__pb2.LookupAttestationsParams.FromString,
          response_serializer=eapi__pb2.LookupAttestationsResponse.SerializeToString,
      ),
      'ResyncPerspectiveGraph': grpc.unary_unary_rpc_method_handler(
          servicer.ResyncPerspectiveGraph,
          request_deserializer=eapi__pb2.ResyncPerspectiveGraphParams.FromString,
          response_serializer=eapi__pb2.ResyncPerspectiveGraphResponse.SerializeToString,
      ),
      'SyncStatus': grpc.unary_unary_rpc_method_handler(
          servicer.SyncStatus,
          request_deserializer=eapi__pb2.SyncParams.FromString,
          response_serializer=eapi__pb2.SyncResponse.SerializeToString,
      ),
      'WaitForSyncComplete': grpc.unary_stream_rpc_method_handler(
          servicer.WaitForSyncComplete,
          request_deserializer=eapi__pb2.SyncParams.FromString,
          response_serializer=eapi__pb2.SyncResponse.SerializeToString,
      ),
      'BuildRTreeProof': grpc.unary_unary_rpc_method_handler(
          servicer.BuildRTreeProof,
          request_deserializer=eapi__pb2.BuildRTreeProofParams.FromString,
          response_serializer=eapi__pb2.BuildRTreeProofResponse.SerializeToString,
      ),
      'VerifyProof': grpc.unary_unary_rpc_method_handler(
          servicer.VerifyProof,
          request_deserializer=eapi__pb2.VerifyProofParams.FromString,
          response_serializer=eapi__pb2.VerifyProofResponse.SerializeToString,
      ),
      'ListLocations': grpc.unary_unary_rpc_method_handler(
          servicer.ListLocations,
          request_deserializer=eapi__pb2.ListLocationsParams.FromString,
          response_serializer=eapi__pb2.ListLocationsResponse.SerializeToString,
      ),
      'Inspect': grpc.unary_unary_rpc_method_handler(
          servicer.Inspect,
          request_deserializer=eapi__pb2.InspectParams.FromString,
          response_serializer=eapi__pb2.InspectResponse.SerializeToString,
      ),
      'ResolveHash': grpc.unary_unary_rpc_method_handler(
          servicer.ResolveHash,
          request_deserializer=eapi__pb2.ResolveHashParams.FromString,
          response_serializer=eapi__pb2.ResolveHashResponse.SerializeToString,
      ),
      'EncryptMessage': grpc.unary_unary_rpc_method_handler(
          servicer.EncryptMessage,
          request_deserializer=eapi__pb2.EncryptMessageParams.FromString,
          response_serializer=eapi__pb2.EncryptMessageResponse.SerializeToString,
      ),
      'DecryptMessage': grpc.unary_unary_rpc_method_handler(
          servicer.DecryptMessage,
          request_deserializer=eapi__pb2.DecryptMessageParams.FromString,
          response_serializer=eapi__pb2.DecryptMessageResponse.SerializeToString,
      ),
      'CreateNameDeclaration': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNameDeclaration,
          request_deserializer=eapi__pb2.CreateNameDeclarationParams.FromString,
          response_serializer=eapi__pb2.CreateNameDeclarationResponse.SerializeToString,
      ),
      'ResolveName': grpc.unary_unary_rpc_method_handler(
          servicer.ResolveName,
          request_deserializer=eapi__pb2.ResolveNameParams.FromString,
          response_serializer=eapi__pb2.ResolveNameResponse.SerializeToString,
      ),
      'MarkEntityInteresting': grpc.unary_unary_rpc_method_handler(
          servicer.MarkEntityInteresting,
          request_deserializer=eapi__pb2.MarkEntityInterestingParams.FromString,
          response_serializer=eapi__pb2.MarkEntityInterestingResponse.SerializeToString,
      ),
      'ResolveReverseName': grpc.unary_unary_rpc_method_handler(
          servicer.ResolveReverseName,
          request_deserializer=eapi__pb2.ResolveReverseNameParams.FromString,
          response_serializer=eapi__pb2.ResolveReverseNameResponse.SerializeToString,
      ),
      'Revoke': grpc.unary_unary_rpc_method_handler(
          servicer.Revoke,
          request_deserializer=eapi__pb2.RevokeParams.FromString,
          response_serializer=eapi__pb2.RevokeResponse.SerializeToString,
      ),
      'CompactProof': grpc.unary_unary_rpc_method_handler(
          servicer.CompactProof,
          request_deserializer=eapi__pb2.CompactProofParams.FromString,
          response_serializer=eapi__pb2.CompactProofResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'mqpb.WAVE', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
