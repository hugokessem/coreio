
CREATE TABLE topup (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    provider provider NOT NULL,
    amount BIGINT NOT NULL,
    is_self BOOLEAN DEFAULT false,
);

CREATE TABLE utility (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    utility_id UUID REFERENCES utility_provider(id) NOT NULL,
    transaction_id UUID REFERENCES transaction(id) NOT NULL,
    metadata JSONB,
);

CREATE TABLE mini_app (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mini_app_id UUID REFERENCES mini_app_provider(id) NOT NULL,
    transaction_id UUID REFERENCES transaction(id) NOT NULL,
    metadata JSONB,
);

CREATE TABLE business (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID REFERENCES business_entity(id) NOT NULL,
    operator_id UUID REFERENCES business_operator(id) NOT NULL,
    transaction_id UUID REFERENCES transaction(id) NOT NULL,
    metat JSONB
);

