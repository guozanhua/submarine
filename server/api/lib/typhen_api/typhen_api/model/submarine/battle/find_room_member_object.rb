# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Battle
  class FindRoomMemberObject
    include Virtus.model(:strict => true)

    attribute :room_member, TyphenApi::Model::Submarine::Battle::RoomMember, :required => false
  end
end
