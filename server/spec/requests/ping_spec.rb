require 'rails_helper'

RSpec.describe "Ping", type: :request do

  describe "POST /ping" do

    context 'with a valid request' do
      before do
        post ping_path, message: 'PING'
      end
      it "should work" do
        expect(response).to have_http_status(200)
      end
      it "should return a message" do
        expect(response_json['message']).to eq 'PING PONG'
      end
    end

    context 'with an invalid request' do
      it "should not work" do
        expect { post ping_path }.to raise_error(Virtus::CoercionError)
      end
    end

  end

end