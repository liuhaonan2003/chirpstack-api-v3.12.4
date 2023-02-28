--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2 (Debian 14.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: application; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    description text NOT NULL,
    organization_id bigint NOT NULL,
    service_profile_id uuid,
    payload_codec text DEFAULT ''::text NOT NULL,
    payload_encoder_script text DEFAULT ''::text NOT NULL,
    payload_decoder_script text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.application OWNER TO postgres;

--
-- Name: application_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.application_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.application_id_seq OWNER TO postgres;

--
-- Name: application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_id_seq OWNED BY public.application.id;


--
-- Name: application_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application_user (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id bigint NOT NULL,
    application_id bigint NOT NULL,
    is_admin boolean NOT NULL
);


ALTER TABLE public.application_user OWNER TO postgres;

--
-- Name: application_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.application_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.application_user_id_seq OWNER TO postgres;

--
-- Name: application_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_user_id_seq OWNED BY public.application_user.id;


--
-- Name: device; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device (
    dev_eui bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    application_id bigint NOT NULL,
    device_profile_id uuid NOT NULL,
    name character varying(100) NOT NULL,
    description text NOT NULL,
    last_seen_at timestamp with time zone,
    device_status_battery integer,
    device_status_margin integer
);


ALTER TABLE public.device OWNER TO postgres;

--
-- Name: device_activation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_activation (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    dev_eui bytea NOT NULL,
    dev_addr bytea NOT NULL,
    app_s_key bytea NOT NULL,
    nwk_s_key bytea NOT NULL
);


ALTER TABLE public.device_activation OWNER TO postgres;

--
-- Name: device_activation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.device_activation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.device_activation_id_seq OWNER TO postgres;

--
-- Name: device_activation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.device_activation_id_seq OWNED BY public.device_activation.id;


--
-- Name: device_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_keys (
    dev_eui bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    app_key bytea NOT NULL,
    join_nonce integer NOT NULL
);


ALTER TABLE public.device_keys OWNER TO postgres;

--
-- Name: device_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_profile (
    device_profile_id uuid NOT NULL,
    network_server_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.device_profile OWNER TO postgres;

--
-- Name: device_queue; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_queue (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    reference character varying(100) NOT NULL,
    dev_eui bytea NOT NULL,
    confirmed boolean DEFAULT false NOT NULL,
    pending boolean DEFAULT false NOT NULL,
    fport smallint NOT NULL,
    data bytea NOT NULL
);


ALTER TABLE public.device_queue OWNER TO postgres;

--
-- Name: device_queue_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.device_queue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.device_queue_id_seq OWNER TO postgres;

--
-- Name: device_queue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.device_queue_id_seq OWNED BY public.device_queue.id;


--
-- Name: device_queue_mapping; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_queue_mapping (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    reference text NOT NULL,
    dev_eui bytea NOT NULL,
    f_cnt integer NOT NULL
);


ALTER TABLE public.device_queue_mapping OWNER TO postgres;

--
-- Name: device_queue_mapping_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.device_queue_mapping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.device_queue_mapping_id_seq OWNER TO postgres;

--
-- Name: device_queue_mapping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.device_queue_mapping_id_seq OWNED BY public.device_queue_mapping.id;


--
-- Name: downlink_queue; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.downlink_queue (
    id bigint NOT NULL,
    reference character varying(100) NOT NULL,
    dev_eui bytea NOT NULL,
    confirmed boolean DEFAULT false NOT NULL,
    pending boolean DEFAULT false NOT NULL,
    fport smallint NOT NULL,
    data bytea NOT NULL
);


ALTER TABLE public.downlink_queue OWNER TO postgres;

--
-- Name: downlink_queue_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.downlink_queue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.downlink_queue_id_seq OWNER TO postgres;

--
-- Name: downlink_queue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.downlink_queue_id_seq OWNED BY public.downlink_queue.id;


--
-- Name: gateway; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gateway (
    mac bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    description text NOT NULL,
    organization_id bigint NOT NULL,
    ping boolean DEFAULT false NOT NULL,
    last_ping_id bigint,
    last_ping_sent_at timestamp with time zone,
    network_server_id bigint
);


ALTER TABLE public.gateway OWNER TO postgres;

--
-- Name: gateway_ping; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gateway_ping (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    gateway_mac bytea NOT NULL,
    frequency integer NOT NULL,
    dr integer NOT NULL
);


ALTER TABLE public.gateway_ping OWNER TO postgres;

--
-- Name: gateway_ping_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gateway_ping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gateway_ping_id_seq OWNER TO postgres;

--
-- Name: gateway_ping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gateway_ping_id_seq OWNED BY public.gateway_ping.id;


--
-- Name: gateway_ping_rx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gateway_ping_rx (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    ping_id bigint NOT NULL,
    gateway_mac bytea NOT NULL,
    received_at timestamp with time zone,
    rssi integer NOT NULL,
    lora_snr numeric(3,1) NOT NULL,
    location point,
    altitude double precision
);


ALTER TABLE public.gateway_ping_rx OWNER TO postgres;

--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gateway_ping_rx_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gateway_ping_rx_id_seq OWNER TO postgres;

--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gateway_ping_rx_id_seq OWNED BY public.gateway_ping_rx.id;


--
-- Name: integration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.integration (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    application_id bigint NOT NULL,
    kind character varying(20) NOT NULL,
    settings jsonb
);


ALTER TABLE public.integration OWNER TO postgres;

--
-- Name: integration_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.integration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.integration_id_seq OWNER TO postgres;

--
-- Name: integration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.integration_id_seq OWNED BY public.integration.id;


--
-- Name: network_server; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.network_server (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    server character varying(255) NOT NULL,
    ca_cert text DEFAULT ''::text NOT NULL,
    tls_cert text DEFAULT ''::text NOT NULL,
    tls_key text DEFAULT ''::text NOT NULL,
    routing_profile_ca_cert text DEFAULT ''::text NOT NULL,
    routing_profile_tls_cert text DEFAULT ''::text NOT NULL,
    routing_profile_tls_key text DEFAULT ''::text NOT NULL,
    gateway_discovery_enabled boolean DEFAULT false NOT NULL,
    gateway_discovery_interval integer DEFAULT 0 NOT NULL,
    gateway_discovery_tx_frequency integer DEFAULT 0 NOT NULL,
    gateway_discovery_dr smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.network_server OWNER TO postgres;

--
-- Name: network_server_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.network_server_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.network_server_id_seq OWNER TO postgres;

--
-- Name: network_server_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.network_server_id_seq OWNED BY public.network_server.id;


--
-- Name: node; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.node (
    dev_eui bytea NOT NULL,
    app_eui bytea NOT NULL,
    app_key bytea NOT NULL,
    used_dev_nonces bytea,
    rx_delay smallint NOT NULL,
    rx1_dr_offset smallint NOT NULL,
    rx_window smallint DEFAULT 0 NOT NULL,
    rx2_dr smallint DEFAULT 0 NOT NULL,
    app_s_key bytea DEFAULT '\x00000000000000000000000000000000'::bytea NOT NULL,
    nwk_s_key bytea DEFAULT '\x00000000000000000000000000000000'::bytea NOT NULL,
    dev_addr bytea DEFAULT '\x00000000'::bytea NOT NULL,
    name character varying(100) DEFAULT ''::character varying NOT NULL,
    relax_fcnt boolean DEFAULT false NOT NULL,
    adr_interval integer DEFAULT 0 NOT NULL,
    installation_margin numeric(5,2) DEFAULT 0 NOT NULL,
    application_id bigint NOT NULL,
    description text NOT NULL,
    is_abp boolean DEFAULT false NOT NULL,
    is_class_c boolean DEFAULT false NOT NULL,
    use_application_settings boolean DEFAULT false NOT NULL
);


ALTER TABLE public.node OWNER TO postgres;

--
-- Name: organization; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organization (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    display_name character varying(100) NOT NULL,
    can_have_gateways boolean NOT NULL
);


ALTER TABLE public.organization OWNER TO postgres;

--
-- Name: organization_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_id_seq OWNER TO postgres;

--
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;


--
-- Name: organization_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organization_user (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    is_admin boolean NOT NULL
);


ALTER TABLE public.organization_user OWNER TO postgres;

--
-- Name: organization_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.organization_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_user_id_seq OWNER TO postgres;

--
-- Name: organization_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.organization_user_id_seq OWNED BY public.organization_user.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: service_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service_profile (
    service_profile_id uuid NOT NULL,
    organization_id bigint NOT NULL,
    network_server_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.service_profile OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    username character varying(100) NOT NULL,
    password_hash character varying(200) NOT NULL,
    session_ttl bigint NOT NULL,
    is_active boolean NOT NULL,
    is_admin boolean NOT NULL,
    email text DEFAULT ''::text NOT NULL,
    note text DEFAULT ''::text NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: application id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application ALTER COLUMN id SET DEFAULT nextval('public.application_id_seq'::regclass);


--
-- Name: application_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_user ALTER COLUMN id SET DEFAULT nextval('public.application_user_id_seq'::regclass);


--
-- Name: device_activation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_activation ALTER COLUMN id SET DEFAULT nextval('public.device_activation_id_seq'::regclass);


--
-- Name: device_queue id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue ALTER COLUMN id SET DEFAULT nextval('public.device_queue_id_seq'::regclass);


--
-- Name: device_queue_mapping id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue_mapping ALTER COLUMN id SET DEFAULT nextval('public.device_queue_mapping_id_seq'::regclass);


--
-- Name: downlink_queue id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.downlink_queue ALTER COLUMN id SET DEFAULT nextval('public.downlink_queue_id_seq'::regclass);


--
-- Name: gateway_ping id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping ALTER COLUMN id SET DEFAULT nextval('public.gateway_ping_id_seq'::regclass);


--
-- Name: gateway_ping_rx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping_rx ALTER COLUMN id SET DEFAULT nextval('public.gateway_ping_rx_id_seq'::regclass);


--
-- Name: integration id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration ALTER COLUMN id SET DEFAULT nextval('public.integration_id_seq'::regclass);


--
-- Name: network_server id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.network_server ALTER COLUMN id SET DEFAULT nextval('public.network_server_id_seq'::regclass);


--
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);


--
-- Name: organization_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_user ALTER COLUMN id SET DEFAULT nextval('public.organization_user_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: application; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.application (id, name, description, organization_id, service_profile_id, payload_codec, payload_encoder_script, payload_decoder_script) FROM stdin;
\.


--
-- Data for Name: application_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.application_user (id, created_at, updated_at, user_id, application_id, is_admin) FROM stdin;
\.


--
-- Data for Name: device; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device (dev_eui, created_at, updated_at, application_id, device_profile_id, name, description, last_seen_at, device_status_battery, device_status_margin) FROM stdin;
\.


--
-- Data for Name: device_activation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device_activation (id, created_at, dev_eui, dev_addr, app_s_key, nwk_s_key) FROM stdin;
\.


--
-- Data for Name: device_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device_keys (dev_eui, created_at, updated_at, app_key, join_nonce) FROM stdin;
\.


--
-- Data for Name: device_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device_profile (device_profile_id, network_server_id, organization_id, created_at, updated_at, name) FROM stdin;
\.


--
-- Data for Name: device_queue; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device_queue (id, created_at, updated_at, reference, dev_eui, confirmed, pending, fport, data) FROM stdin;
\.


--
-- Data for Name: device_queue_mapping; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.device_queue_mapping (id, created_at, reference, dev_eui, f_cnt) FROM stdin;
\.


--
-- Data for Name: downlink_queue; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.downlink_queue (id, reference, dev_eui, confirmed, pending, fport, data) FROM stdin;
\.


--
-- Data for Name: gateway; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.gateway (mac, created_at, updated_at, name, description, organization_id, ping, last_ping_id, last_ping_sent_at, network_server_id) FROM stdin;
\.


--
-- Data for Name: gateway_ping; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.gateway_ping (id, created_at, gateway_mac, frequency, dr) FROM stdin;
\.


--
-- Data for Name: gateway_ping_rx; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.gateway_ping_rx (id, created_at, ping_id, gateway_mac, received_at, rssi, lora_snr, location, altitude) FROM stdin;
\.


--
-- Data for Name: integration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.integration (id, created_at, updated_at, application_id, kind, settings) FROM stdin;
\.


--
-- Data for Name: network_server; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.network_server (id, created_at, updated_at, name, server, ca_cert, tls_cert, tls_key, routing_profile_ca_cert, routing_profile_tls_cert, routing_profile_tls_key, gateway_discovery_enabled, gateway_discovery_interval, gateway_discovery_tx_frequency, gateway_discovery_dr) FROM stdin;
\.


--
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.node (dev_eui, app_eui, app_key, used_dev_nonces, rx_delay, rx1_dr_offset, rx_window, rx2_dr, app_s_key, nwk_s_key, dev_addr, name, relax_fcnt, adr_interval, installation_margin, application_id, description, is_abp, is_class_c, use_application_settings) FROM stdin;
\.


--
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organization (id, created_at, updated_at, name, display_name, can_have_gateways) FROM stdin;
1	2022-06-16 11:19:17.714778+00	2022-06-16 11:19:17.714778+00	chirpstack	ChirpStack	t
\.


--
-- Data for Name: organization_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organization_user (id, created_at, updated_at, user_id, organization_id, is_admin) FROM stdin;
1	2022-06-16 11:19:17.714778+00	2022-06-16 11:19:17.714778+00	1	1	t
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
27	t
\.


--
-- Data for Name: service_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.service_profile (service_profile_id, organization_id, network_server_id, created_at, updated_at, name) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, created_at, updated_at, username, password_hash, session_ttl, is_active, is_admin, email, note) FROM stdin;
1	2022-06-16 11:19:17.670204+00	2022-06-16 11:19:17.670204+00	admin	PBKDF2$sha512$1$l8zGKtxRESq3PA2kFhHRWA==$H3lGMxOt55wjwoc+myeOoABofJY9oDpldJa7fhqdjbh700V6FLPML75UmBOt9J5VFNjAL1AvqCozA1HJM0QVGA==	0	t	t	admin	
\.


--
-- Name: application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.application_id_seq', 1, false);


--
-- Name: application_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.application_user_id_seq', 1, false);


--
-- Name: device_activation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_activation_id_seq', 1, false);


--
-- Name: device_queue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_queue_id_seq', 1, false);


--
-- Name: device_queue_mapping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_queue_mapping_id_seq', 1, false);


--
-- Name: downlink_queue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.downlink_queue_id_seq', 1, false);


--
-- Name: gateway_ping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gateway_ping_id_seq', 1, false);


--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gateway_ping_rx_id_seq', 1, false);


--
-- Name: integration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.integration_id_seq', 1, false);


--
-- Name: network_server_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.network_server_id_seq', 1, false);


--
-- Name: organization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.organization_id_seq', 1, true);


--
-- Name: organization_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.organization_user_id_seq', 1, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- Name: application application_name_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_name_organization_id_key UNIQUE (name, organization_id);


--
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- Name: application_user application_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_user
    ADD CONSTRAINT application_user_pkey PRIMARY KEY (id);


--
-- Name: application_user application_user_user_id_application_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_user
    ADD CONSTRAINT application_user_user_id_application_id_key UNIQUE (user_id, application_id);


--
-- Name: device_activation device_activation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_activation
    ADD CONSTRAINT device_activation_pkey PRIMARY KEY (id);


--
-- Name: device_keys device_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_keys
    ADD CONSTRAINT device_keys_pkey PRIMARY KEY (dev_eui);


--
-- Name: device device_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_pkey PRIMARY KEY (dev_eui);


--
-- Name: device_profile device_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_pkey PRIMARY KEY (device_profile_id);


--
-- Name: device_queue_mapping device_queue_mapping_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue_mapping
    ADD CONSTRAINT device_queue_mapping_pkey PRIMARY KEY (id);


--
-- Name: device_queue device_queue_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue
    ADD CONSTRAINT device_queue_pkey PRIMARY KEY (id);


--
-- Name: gateway gateway_name_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_name_organization_id_key UNIQUE (name, organization_id);


--
-- Name: gateway_ping gateway_ping_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping
    ADD CONSTRAINT gateway_ping_pkey PRIMARY KEY (id);


--
-- Name: gateway_ping_rx gateway_ping_rx_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_pkey PRIMARY KEY (id);


--
-- Name: gateway gateway_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_pkey PRIMARY KEY (mac);


--
-- Name: integration integration_kind_application_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_kind_application_id UNIQUE (kind, application_id);


--
-- Name: integration integration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_pkey PRIMARY KEY (id);


--
-- Name: network_server network_server_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.network_server
    ADD CONSTRAINT network_server_pkey PRIMARY KEY (id);


--
-- Name: node node_application_id_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_application_id_name_key UNIQUE (application_id, name);


--
-- Name: node node_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_pkey PRIMARY KEY (dev_eui);


--
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: organization_user organization_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_pkey PRIMARY KEY (id);


--
-- Name: organization_user organization_user_user_id_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_user_id_organization_id_key UNIQUE (user_id, organization_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: service_profile service_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_pkey PRIMARY KEY (service_profile_id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: device_queue_mapping_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX device_queue_mapping_created_at ON public.device_queue_mapping USING btree (created_at);


--
-- Name: device_queue_mapping_dev_eui; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX device_queue_mapping_dev_eui ON public.device_queue_mapping USING btree (dev_eui);


--
-- Name: downlink_queue_dev_eui; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX downlink_queue_dev_eui ON public.downlink_queue USING btree (dev_eui);


--
-- Name: idx_application_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_application_name ON public.application USING btree (name);


--
-- Name: idx_application_organization_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_application_organization_id ON public.application USING btree (organization_id);


--
-- Name: idx_application_service_profile_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_application_service_profile_id ON public.application USING btree (service_profile_id);


--
-- Name: idx_application_user_application_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_application_user_application_id ON public.application_user USING btree (application_id);


--
-- Name: idx_application_user_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_application_user_user_id ON public.application_user USING btree (user_id);


--
-- Name: idx_device_activation_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_activation_created_at ON public.device_activation USING btree (created_at);


--
-- Name: idx_device_activation_dev_eui; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_activation_dev_eui ON public.device_activation USING btree (dev_eui);


--
-- Name: idx_device_application_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_application_id ON public.device USING btree (application_id);


--
-- Name: idx_device_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_created_at ON public.device USING btree (created_at);


--
-- Name: idx_device_dev_eui_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_dev_eui_prefix ON public.device USING btree (encode(dev_eui, 'hex'::text) varchar_pattern_ops);


--
-- Name: idx_device_device_profile_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_device_profile_id ON public.device USING btree (device_profile_id);


--
-- Name: idx_device_keys_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_keys_created_at ON public.device_keys USING btree (created_at);


--
-- Name: idx_device_keys_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_keys_updated_at ON public.device_keys USING btree (updated_at);


--
-- Name: idx_device_name_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_name_prefix ON public.device USING btree (name varchar_pattern_ops);


--
-- Name: idx_device_profile_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_profile_created_at ON public.device_profile USING btree (created_at);


--
-- Name: idx_device_profile_network_server_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_profile_network_server_id ON public.device_profile USING btree (network_server_id);


--
-- Name: idx_device_profile_organization_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_profile_organization_id ON public.device_profile USING btree (organization_id);


--
-- Name: idx_device_profile_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_profile_updated_at ON public.device_profile USING btree (updated_at);


--
-- Name: idx_device_queue_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_queue_created_at ON public.device_queue USING btree (created_at);


--
-- Name: idx_device_queue_dev_eui; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_queue_dev_eui ON public.device_queue USING btree (dev_eui);


--
-- Name: idx_device_queue_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_queue_updated_at ON public.device_queue USING btree (updated_at);


--
-- Name: idx_device_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_device_updated_at ON public.device USING btree (updated_at);


--
-- Name: idx_gateway_last_ping_sent_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_last_ping_sent_at ON public.gateway USING btree (last_ping_sent_at);


--
-- Name: idx_gateway_network_server_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_network_server_id ON public.gateway USING btree (network_server_id);


--
-- Name: idx_gateway_organization_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_organization_id ON public.gateway USING btree (organization_id);


--
-- Name: idx_gateway_ping; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping ON public.gateway USING btree (ping);


--
-- Name: idx_gateway_ping_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping_created_at ON public.gateway_ping USING btree (created_at);


--
-- Name: idx_gateway_ping_gateway_mac; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping_gateway_mac ON public.gateway_ping USING btree (gateway_mac);


--
-- Name: idx_gateway_ping_rx_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping_rx_created_at ON public.gateway_ping_rx USING btree (created_at);


--
-- Name: idx_gateway_ping_rx_gateway_mac; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping_rx_gateway_mac ON public.gateway_ping_rx USING btree (gateway_mac);


--
-- Name: idx_gateway_ping_rx_ping_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_gateway_ping_rx_ping_id ON public.gateway_ping_rx USING btree (ping_id);


--
-- Name: idx_integration_application_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_integration_application_id ON public.integration USING btree (application_id);


--
-- Name: idx_integration_kind; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_integration_kind ON public.integration USING btree (kind);


--
-- Name: idx_network_server_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_network_server_created_at ON public.network_server USING btree (created_at);


--
-- Name: idx_network_server_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_network_server_updated_at ON public.network_server USING btree (updated_at);


--
-- Name: idx_node_application_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_node_application_id ON public.node USING btree (application_id);


--
-- Name: idx_node_dev_eui_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_node_dev_eui_prefix ON public.node USING btree (encode(dev_eui, 'hex'::text) varchar_pattern_ops);


--
-- Name: idx_node_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_node_name ON public.node USING btree (name);


--
-- Name: idx_node_name_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_node_name_prefix ON public.node USING btree (name varchar_pattern_ops);


--
-- Name: idx_organization_display_name_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_organization_display_name_prefix ON public.organization USING btree (lower((display_name)::text) varchar_pattern_ops);


--
-- Name: idx_organization_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_organization_name ON public.organization USING btree (name);


--
-- Name: idx_organization_user_organization_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_organization_user_organization_id ON public.organization_user USING btree (organization_id);


--
-- Name: idx_organization_user_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_organization_user_user_id ON public.organization_user USING btree (user_id);


--
-- Name: idx_service_profile_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_service_profile_created_at ON public.service_profile USING btree (created_at);


--
-- Name: idx_service_profile_network_server_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_service_profile_network_server_id ON public.service_profile USING btree (network_server_id);


--
-- Name: idx_service_profile_organization_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_service_profile_organization_id ON public.service_profile USING btree (organization_id);


--
-- Name: idx_service_profile_updated_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_service_profile_updated_at ON public.service_profile USING btree (updated_at);


--
-- Name: idx_user_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_username ON public."user" USING btree (username);


--
-- Name: idx_user_username_prefix; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_user_username_prefix ON public."user" USING btree (username varchar_pattern_ops);


--
-- Name: node_app_eui; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX node_app_eui ON public.node USING btree (app_eui);


--
-- Name: application application_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: application application_service_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_service_profile_id_fkey FOREIGN KEY (service_profile_id) REFERENCES public.service_profile(service_profile_id);


--
-- Name: application_user application_user_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_user
    ADD CONSTRAINT application_user_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: application_user application_user_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_user
    ADD CONSTRAINT application_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: device_activation device_activation_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_activation
    ADD CONSTRAINT device_activation_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: device device_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id);


--
-- Name: device device_device_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_device_profile_id_fkey FOREIGN KEY (device_profile_id) REFERENCES public.device_profile(device_profile_id);


--
-- Name: device_keys device_keys_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_keys
    ADD CONSTRAINT device_keys_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: device_profile device_profile_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: device_profile device_profile_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);


--
-- Name: device_queue device_queue_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue
    ADD CONSTRAINT device_queue_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: device_queue_mapping device_queue_mapping_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_queue_mapping
    ADD CONSTRAINT device_queue_mapping_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: downlink_queue downlink_queue_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.downlink_queue
    ADD CONSTRAINT downlink_queue_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.node(dev_eui) ON DELETE CASCADE;


--
-- Name: gateway gateway_last_ping_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_last_ping_id_fkey FOREIGN KEY (last_ping_id) REFERENCES public.gateway_ping(id) ON DELETE SET NULL;


--
-- Name: gateway gateway_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: gateway gateway_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: gateway_ping gateway_ping_gateway_mac_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping
    ADD CONSTRAINT gateway_ping_gateway_mac_fkey FOREIGN KEY (gateway_mac) REFERENCES public.gateway(mac) ON DELETE CASCADE;


--
-- Name: gateway_ping_rx gateway_ping_rx_gateway_mac_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_gateway_mac_fkey FOREIGN KEY (gateway_mac) REFERENCES public.gateway(mac) ON DELETE CASCADE;


--
-- Name: gateway_ping_rx gateway_ping_rx_ping_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_ping_id_fkey FOREIGN KEY (ping_id) REFERENCES public.gateway_ping(id) ON DELETE CASCADE;


--
-- Name: integration integration_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: node node_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: organization_user organization_user_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: organization_user organization_user_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: service_profile service_profile_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: service_profile service_profile_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);


--
-- PostgreSQL database dump complete
--

