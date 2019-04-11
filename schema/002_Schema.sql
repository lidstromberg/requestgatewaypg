CREATE TABLE public.gateway
(
    gatewayid bigserial not null,
    appscope character varying(510) COLLATE pg_catalog."default" NOT NULL,
    remoteaddress character varying(255) COLLATE pg_catalog."default" NOT NULL,
    createddate timestamp with time zone NOT NULL DEFAULT now(),
    retireddate timestamp with time zone,
    lasttouched timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT pk_gateway PRIMARY KEY (gatewayid),
    CONSTRAINT uc_gateway_1 UNIQUE (appscope,remoteaddress)
);

ALTER TABLE public.gateway OWNER to postgres;

GRANT ALL ON TABLE public.gateway to gatewayuser;
GRANT ALL ON SEQUENCE gateway_gatewayid_seq to gatewayuser;

