module TyphenApiRespondable
  RENDER_FORMAT = :msgpack

  def params
    @typhen_api_params ||= self.class::RequestType.new(super.except(:controller, :action).to_unsafe_h)
  end

  def render(response_body, error: false)
    if error
      super RENDER_FORMAT => TyphenApi.to_serializable_object(self.class::ErrorType.new(response_body)), status: 500
    elsif self.class::ResponseType.present?
      super RENDER_FORMAT => TyphenApi.to_serializable_object(self.class::ResponseType.new(response_body))
    else
      super RENDER_FORMAT => nil
    end
  end
end
