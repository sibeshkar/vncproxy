# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: addressbook.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='addressbook.proto',
  package='protofile',
  syntax='proto3',
  serialized_options=_b('\n\024com.example.tutorialB\021AddressBookProtos\252\002$Google.Protobuf.Examples.AddressBook'),
  serialized_pb=_b('\n\x11\x61\x64\x64ressbook.proto\x12\tprotofile\x1a\x1fgoogle/protobuf/timestamp.proto\"\x89\x02\n\x06Person\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\x05\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12-\n\x06phones\x18\x04 \x03(\x0b\x32\x1d.protofile.Person.PhoneNumber\x12\x30\n\x0clast_updated\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x1aH\n\x0bPhoneNumber\x12\x0e\n\x06number\x18\x01 \x01(\t\x12)\n\x04type\x18\x02 \x01(\x0e\x32\x1b.protofile.Person.PhoneType\"+\n\tPhoneType\x12\n\n\x06MOBILE\x10\x00\x12\x08\n\x04HOME\x10\x01\x12\x08\n\x04WORK\x10\x02\"0\n\x0b\x41\x64\x64ressBook\x12!\n\x06people\x18\x01 \x03(\x0b\x32\x11.protofile.PersonBP\n\x14\x63om.example.tutorialB\x11\x41\x64\x64ressBookProtos\xaa\x02$Google.Protobuf.Examples.AddressBookb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])



_PERSON_PHONETYPE = _descriptor.EnumDescriptor(
  name='PhoneType',
  full_name='protofile.Person.PhoneType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='MOBILE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HOME', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WORK', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=288,
  serialized_end=331,
)
_sym_db.RegisterEnumDescriptor(_PERSON_PHONETYPE)


_PERSON_PHONENUMBER = _descriptor.Descriptor(
  name='PhoneNumber',
  full_name='protofile.Person.PhoneNumber',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='number', full_name='protofile.Person.PhoneNumber.number', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='protofile.Person.PhoneNumber.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=214,
  serialized_end=286,
)

_PERSON = _descriptor.Descriptor(
  name='Person',
  full_name='protofile.Person',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='protofile.Person.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id', full_name='protofile.Person.id', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='protofile.Person.email', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phones', full_name='protofile.Person.phones', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='last_updated', full_name='protofile.Person.last_updated', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_PERSON_PHONENUMBER, ],
  enum_types=[
    _PERSON_PHONETYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=66,
  serialized_end=331,
)


_ADDRESSBOOK = _descriptor.Descriptor(
  name='AddressBook',
  full_name='protofile.AddressBook',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='people', full_name='protofile.AddressBook.people', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=333,
  serialized_end=381,
)

_PERSON_PHONENUMBER.fields_by_name['type'].enum_type = _PERSON_PHONETYPE
_PERSON_PHONENUMBER.containing_type = _PERSON
_PERSON.fields_by_name['phones'].message_type = _PERSON_PHONENUMBER
_PERSON.fields_by_name['last_updated'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_PERSON_PHONETYPE.containing_type = _PERSON
_ADDRESSBOOK.fields_by_name['people'].message_type = _PERSON
DESCRIPTOR.message_types_by_name['Person'] = _PERSON
DESCRIPTOR.message_types_by_name['AddressBook'] = _ADDRESSBOOK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Person = _reflection.GeneratedProtocolMessageType('Person', (_message.Message,), dict(

  PhoneNumber = _reflection.GeneratedProtocolMessageType('PhoneNumber', (_message.Message,), dict(
    DESCRIPTOR = _PERSON_PHONENUMBER,
    __module__ = 'addressbook_pb2'
    # @@protoc_insertion_point(class_scope:protofile.Person.PhoneNumber)
    ))
  ,
  DESCRIPTOR = _PERSON,
  __module__ = 'addressbook_pb2'
  # @@protoc_insertion_point(class_scope:protofile.Person)
  ))
_sym_db.RegisterMessage(Person)
_sym_db.RegisterMessage(Person.PhoneNumber)

AddressBook = _reflection.GeneratedProtocolMessageType('AddressBook', (_message.Message,), dict(
  DESCRIPTOR = _ADDRESSBOOK,
  __module__ = 'addressbook_pb2'
  # @@protoc_insertion_point(class_scope:protofile.AddressBook)
  ))
_sym_db.RegisterMessage(AddressBook)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
