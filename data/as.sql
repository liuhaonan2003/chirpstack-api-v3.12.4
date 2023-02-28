--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.24
-- Dumped by pg_dump version 9.6.24

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

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: hstore; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS hstore WITH SCHEMA public;


--
-- Name: EXTENSION hstore; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION hstore IS 'data type for storing sets of (key, value) pairs';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: api_key; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.api_key (
    id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    organization_id bigint,
    application_id bigint
);


ALTER TABLE public.api_key OWNER TO chirpstack_as;

--
-- Name: application; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.application (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    description text NOT NULL,
    organization_id bigint NOT NULL,
    service_profile_id uuid NOT NULL,
    payload_codec text DEFAULT ''::text NOT NULL,
    payload_encoder_script text DEFAULT ''::text NOT NULL,
    payload_decoder_script text DEFAULT ''::text NOT NULL,
    mqtt_tls_cert bytea
);


ALTER TABLE public.application OWNER TO chirpstack_as;

--
-- Name: application_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.application_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.application_id_seq OWNER TO chirpstack_as;

--
-- Name: application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.application_id_seq OWNED BY public.application.id;


--
-- Name: code_migration; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.code_migration (
    id text NOT NULL,
    applied_at timestamp with time zone NOT NULL
);


ALTER TABLE public.code_migration OWNER TO chirpstack_as;

--
-- Name: device; Type: TABLE; Schema: public; Owner: chirpstack_as
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
    device_status_battery numeric(5,2),
    device_status_margin integer,
    latitude double precision,
    longitude double precision,
    altitude double precision,
    device_status_external_power_source boolean NOT NULL,
    dr smallint,
    variables public.hstore,
    tags public.hstore,
    dev_addr bytea NOT NULL,
    app_s_key bytea NOT NULL
);


ALTER TABLE public.device OWNER TO chirpstack_as;

--
-- Name: device_keys; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.device_keys (
    dev_eui bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    nwk_key bytea NOT NULL,
    join_nonce integer NOT NULL,
    app_key bytea NOT NULL
);


ALTER TABLE public.device_keys OWNER TO chirpstack_as;

--
-- Name: device_multicast_group; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.device_multicast_group (
    dev_eui bytea NOT NULL,
    multicast_group_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.device_multicast_group OWNER TO chirpstack_as;

--
-- Name: device_profile; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.device_profile (
    device_profile_id uuid NOT NULL,
    network_server_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    payload_codec text NOT NULL,
    payload_encoder_script text NOT NULL,
    payload_decoder_script text NOT NULL,
    tags public.hstore,
    uplink_interval bigint NOT NULL
);


ALTER TABLE public.device_profile OWNER TO chirpstack_as;

--
-- Name: gateway; Type: TABLE; Schema: public; Owner: chirpstack_as
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
    network_server_id bigint NOT NULL,
    gateway_profile_id uuid,
    first_seen_at timestamp with time zone,
    last_seen_at timestamp with time zone,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    altitude double precision NOT NULL,
    tags public.hstore,
    metadata public.hstore,
    service_profile_id uuid
);


ALTER TABLE public.gateway OWNER TO chirpstack_as;

--
-- Name: gateway_ping; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.gateway_ping (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    gateway_mac bytea NOT NULL,
    frequency integer NOT NULL,
    dr integer NOT NULL
);


ALTER TABLE public.gateway_ping OWNER TO chirpstack_as;

--
-- Name: gateway_ping_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.gateway_ping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gateway_ping_id_seq OWNER TO chirpstack_as;

--
-- Name: gateway_ping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.gateway_ping_id_seq OWNED BY public.gateway_ping.id;


--
-- Name: gateway_ping_rx; Type: TABLE; Schema: public; Owner: chirpstack_as
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


ALTER TABLE public.gateway_ping_rx OWNER TO chirpstack_as;

--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.gateway_ping_rx_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gateway_ping_rx_id_seq OWNER TO chirpstack_as;

--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.gateway_ping_rx_id_seq OWNED BY public.gateway_ping_rx.id;


--
-- Name: gateway_profile; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.gateway_profile (
    gateway_profile_id uuid NOT NULL,
    network_server_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    stats_interval bigint NOT NULL
);


ALTER TABLE public.gateway_profile OWNER TO chirpstack_as;

--
-- Name: integration; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.integration (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    application_id bigint NOT NULL,
    kind character varying(20) NOT NULL,
    settings jsonb
);


ALTER TABLE public.integration OWNER TO chirpstack_as;

--
-- Name: integration_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.integration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.integration_id_seq OWNER TO chirpstack_as;

--
-- Name: integration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.integration_id_seq OWNED BY public.integration.id;


--
-- Name: multicast_group; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.multicast_group (
    id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    mc_app_s_key bytea,
    application_id bigint NOT NULL
);


ALTER TABLE public.multicast_group OWNER TO chirpstack_as;

--
-- Name: network_server; Type: TABLE; Schema: public; Owner: chirpstack_as
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


ALTER TABLE public.network_server OWNER TO chirpstack_as;

--
-- Name: network_server_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.network_server_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.network_server_id_seq OWNER TO chirpstack_as;

--
-- Name: network_server_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.network_server_id_seq OWNED BY public.network_server.id;


--
-- Name: organization; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.organization (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL,
    display_name character varying(100) NOT NULL,
    can_have_gateways boolean NOT NULL,
    max_device_count integer NOT NULL,
    max_gateway_count integer NOT NULL
);


ALTER TABLE public.organization OWNER TO chirpstack_as;

--
-- Name: organization_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_id_seq OWNER TO chirpstack_as;

--
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;


--
-- Name: organization_user; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.organization_user (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    is_admin boolean NOT NULL,
    is_device_admin boolean NOT NULL,
    is_gateway_admin boolean NOT NULL
);


ALTER TABLE public.organization_user OWNER TO chirpstack_as;

--
-- Name: organization_user_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.organization_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_user_id_seq OWNER TO chirpstack_as;

--
-- Name: organization_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.organization_user_id_seq OWNED BY public.organization_user.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO chirpstack_as;

--
-- Name: service_profile; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public.service_profile (
    service_profile_id uuid NOT NULL,
    organization_id bigint NOT NULL,
    network_server_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.service_profile OWNER TO chirpstack_as;

--
-- Name: user; Type: TABLE; Schema: public; Owner: chirpstack_as
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    email text NOT NULL,
    password_hash character varying(200) NOT NULL,
    session_ttl bigint NOT NULL,
    is_active boolean NOT NULL,
    is_admin boolean NOT NULL,
    email_old text DEFAULT ''::text NOT NULL,
    note text NOT NULL,
    external_id text,
    email_verified boolean NOT NULL
);


ALTER TABLE public."user" OWNER TO chirpstack_as;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: chirpstack_as
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO chirpstack_as;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: chirpstack_as
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: application id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.application ALTER COLUMN id SET DEFAULT nextval('public.application_id_seq'::regclass);


--
-- Name: gateway_ping id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping ALTER COLUMN id SET DEFAULT nextval('public.gateway_ping_id_seq'::regclass);


--
-- Name: gateway_ping_rx id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping_rx ALTER COLUMN id SET DEFAULT nextval('public.gateway_ping_rx_id_seq'::regclass);


--
-- Name: integration id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.integration ALTER COLUMN id SET DEFAULT nextval('public.integration_id_seq'::regclass);


--
-- Name: network_server id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.network_server ALTER COLUMN id SET DEFAULT nextval('public.network_server_id_seq'::regclass);


--
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);


--
-- Name: organization_user id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization_user ALTER COLUMN id SET DEFAULT nextval('public.organization_user_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: api_key; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.api_key (id, created_at, name, is_admin, organization_id, application_id) FROM stdin;
\.


--
-- Data for Name: application; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.application (id, name, description, organization_id, service_profile_id, payload_codec, payload_encoder_script, payload_decoder_script, mqtt_tls_cert) FROM stdin;
3	loraserver-app	loraserver-app-test	2	098690a8-11e2-4c8c-a99f-12b769bbcab5				\\x
\.


--
-- Name: application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.application_id_seq', 3, true);


--
-- Data for Name: code_migration; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.code_migration (id, applied_at) FROM stdin;
migrate_gw_stats	2022-06-14 01:28:04.546736+00
migrate_to_cluster_keys	2022-06-14 01:28:04.549412+00
migrate_to_golang_migrate	2022-06-15 11:34:48.163072+00
validate_multicast_group_devices	2022-06-15 11:34:48.169976+00
\.


--
-- Data for Name: device; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.device (dev_eui, created_at, updated_at, application_id, device_profile_id, name, description, last_seen_at, device_status_battery, device_status_margin, latitude, longitude, altitude, device_status_external_power_source, dr, variables, tags, dev_addr, app_s_key) FROM stdin;
\\xffffff10000340b2	2022-06-14 09:37:38.037969+00	2022-06-16 06:28:19.925343+00	3	a816abbd-3fac-41a8-be5f-bb7787a3eb2a	0000220110000004	0000220110000004	\N	\N	\N	\N	\N	\N	f	\N			\\x068e7784	\\xe0418240a16429c85900800204e9832b
\.


--
-- Data for Name: device_keys; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.device_keys (dev_eui, created_at, updated_at, nwk_key, join_nonce, app_key) FROM stdin;
\\xffffff10000340b2	2022-06-16 05:44:43.398077+00	2022-06-22 02:01:32.630568+00	\\x2b7e151628aed2a6abf7158809cf4f3c	259	\\x00000000000000000000000000000000
\.


--
-- Data for Name: device_multicast_group; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.device_multicast_group (dev_eui, multicast_group_id, created_at) FROM stdin;
\\xffffff10000340b2	a50d31cd-0053-46c4-bf4e-48f892e28f3c	2022-06-15 00:35:44.749319+00
\.


--
-- Data for Name: device_profile; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.device_profile (device_profile_id, network_server_id, organization_id, created_at, updated_at, name, payload_codec, payload_encoder_script, payload_decoder_script, tags, uplink_interval) FROM stdin;
b295f31b-a07e-4141-bd29-1769a2daaf86	17	2	2022-06-14 02:58:13.96097+00	2022-06-15 02:04:37.740125+00	dp-cn470-abp-classA					600000000000
2781bb4b-2446-4426-b95b-75a088835b87	17	2	2022-06-14 02:59:41.375083+00	2022-06-15 02:04:51.336093+00	dp-cn470-abp-classC					900000000000
c7368184-bdf7-47c6-a4ef-18c038941f2e	17	2	2022-06-14 03:15:01.495264+00	2022-06-15 02:04:57.28853+00	dp-cn470-abp-classD					9000000000000
a12dff5e-7870-4f3a-9b5b-334c3c4ff8ae	17	2	2022-06-14 03:15:59.935255+00	2022-06-15 02:05:09.128494+00	dp-cn470-otaa_classA					600000000000
a816abbd-3fac-41a8-be5f-bb7787a3eb2a	17	2	2022-06-14 03:16:46.690533+00	2022-06-15 02:05:14.863915+00	dp-cn470-otaa_classC					900000000000
\.


--
-- Data for Name: gateway; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.gateway (mac, created_at, updated_at, name, description, organization_id, ping, last_ping_id, last_ping_sent_at, network_server_id, gateway_profile_id, first_seen_at, last_seen_at, latitude, longitude, altitude, tags, metadata, service_profile_id) FROM stdin;
\\x0005f041c8e3dc84	2022-06-14 08:29:19.096368+00	2022-06-22 02:50:41.422387+00	0005f041c8e3dc84	0005f041c8e3dc84	2	f	\N	\N	17	9533cdd4-219e-4a65-89d0-cccf9e0680ef	2022-06-14 08:29:47.249269+00	2022-06-22 02:50:41.421753+00	22.5672200309229503	114.165715988196425	0			098690a8-11e2-4c8c-a99f-12b769bbcab5
\\xb827ebfffe7d039f	2022-06-14 08:29:50.47555+00	2022-06-14 08:29:50.47555+00	b827ebfffe7d039f	b827ebfffe7d039f	2	f	\N	\N	17	9533cdd4-219e-4a65-89d0-cccf9e0680ef	\N	\N	22.5672211367136164	114.165713634916429	0		\N	098690a8-11e2-4c8c-a99f-12b769bbcab5
\.


--
-- Data for Name: gateway_ping; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.gateway_ping (id, created_at, gateway_mac, frequency, dr) FROM stdin;
\.


--
-- Name: gateway_ping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.gateway_ping_id_seq', 1, false);


--
-- Data for Name: gateway_ping_rx; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.gateway_ping_rx (id, created_at, ping_id, gateway_mac, received_at, rssi, lora_snr, location, altitude) FROM stdin;
\.


--
-- Name: gateway_ping_rx_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.gateway_ping_rx_id_seq', 1, false);


--
-- Data for Name: gateway_profile; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.gateway_profile (gateway_profile_id, network_server_id, created_at, updated_at, name, stats_interval) FROM stdin;
9533cdd4-219e-4a65-89d0-cccf9e0680ef	17	2022-06-14 02:52:41.597983+00	2022-06-15 02:19:16.724771+00	gp-cn470	120000000000
\.


--
-- Data for Name: integration; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.integration (id, created_at, updated_at, application_id, kind, settings) FROM stdin;
\.


--
-- Name: integration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.integration_id_seq', 1, false);


--
-- Data for Name: multicast_group; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.multicast_group (id, created_at, updated_at, name, mc_app_s_key, application_id) FROM stdin;
a50d31cd-0053-46c4-bf4e-48f892e28f3c	2022-06-15 00:34:40.61292+00	2022-06-15 00:34:40.61292+00	mg	\\x22222222222222222222222222222222	3
\.


--
-- Data for Name: network_server; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.network_server (id, created_at, updated_at, name, server, ca_cert, tls_cert, tls_key, routing_profile_ca_cert, routing_profile_tls_cert, routing_profile_tls_key, gateway_discovery_enabled, gateway_discovery_interval, gateway_discovery_tx_frequency, gateway_discovery_dr) FROM stdin;
17	2022-06-14 02:50:56.00763+00	2022-06-14 05:50:36.627557+00	NS-CN470	chirpstack-network-server:8000							f	0	0	0
\.


--
-- Name: network_server_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.network_server_id_seq', 22, true);


--
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.organization (id, created_at, updated_at, name, display_name, can_have_gateways, max_device_count, max_gateway_count) FROM stdin;
1	2022-06-14 01:28:02.990029+00	2022-06-14 01:28:02.990029+00	chirpstack	ChirpStack	t	0	0
2	2022-06-14 02:37:54.531709+00	2022-06-15 02:15:37.124113+00	loraserver	LoRa Server	t	0	0
\.


--
-- Name: organization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.organization_id_seq', 2, true);


--
-- Data for Name: organization_user; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.organization_user (id, created_at, updated_at, user_id, organization_id, is_admin, is_device_admin, is_gateway_admin) FROM stdin;
1	2022-06-14 01:28:02.990029+00	2022-06-14 01:28:02.990029+00	1	1	t	f	f
\.


--
-- Name: organization_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.organization_user_id_seq', 1, true);


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.schema_migrations (version, dirty) FROM stdin;
60	f
\.


--
-- Data for Name: service_profile; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public.service_profile (service_profile_id, organization_id, network_server_id, created_at, updated_at, name) FROM stdin;
098690a8-11e2-4c8c-a99f-12b769bbcab5	2	17	2022-06-14 02:53:44.187887+00	2022-06-14 02:53:44.187887+00	sp-cn470
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: chirpstack_as
--

COPY public."user" (id, created_at, updated_at, email, password_hash, session_ttl, is_active, is_admin, email_old, note, external_id, email_verified) FROM stdin;
1	2022-06-14 01:28:02.944157+00	2022-06-14 01:28:02.944157+00	admin	PBKDF2$sha512$1$l8zGKtxRESq3PA2kFhHRWA==$H3lGMxOt55wjwoc+myeOoABofJY9oDpldJa7fhqdjbh700V6FLPML75UmBOt9J5VFNjAL1AvqCozA1HJM0QVGA==	0	t	t			\N	f
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: chirpstack_as
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- Name: api_key api_key_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.api_key
    ADD CONSTRAINT api_key_pkey PRIMARY KEY (id);


--
-- Name: application application_name_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_name_organization_id_key UNIQUE (name, organization_id);


--
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- Name: code_migration code_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.code_migration
    ADD CONSTRAINT code_migration_pkey PRIMARY KEY (id);


--
-- Name: device_keys device_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_keys
    ADD CONSTRAINT device_keys_pkey PRIMARY KEY (dev_eui);


--
-- Name: device_multicast_group device_multicast_group_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_multicast_group
    ADD CONSTRAINT device_multicast_group_pkey PRIMARY KEY (multicast_group_id, dev_eui);


--
-- Name: device device_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_pkey PRIMARY KEY (dev_eui);


--
-- Name: device_profile device_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_pkey PRIMARY KEY (device_profile_id);


--
-- Name: gateway gateway_name_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_name_organization_id_key UNIQUE (name, organization_id);


--
-- Name: gateway_ping gateway_ping_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping
    ADD CONSTRAINT gateway_ping_pkey PRIMARY KEY (id);


--
-- Name: gateway_ping_rx gateway_ping_rx_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_pkey PRIMARY KEY (id);


--
-- Name: gateway gateway_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_pkey PRIMARY KEY (mac);


--
-- Name: gateway_profile gateway_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_profile
    ADD CONSTRAINT gateway_profile_pkey PRIMARY KEY (gateway_profile_id);


--
-- Name: integration integration_kind_application_id; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_kind_application_id UNIQUE (kind, application_id);


--
-- Name: integration integration_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_pkey PRIMARY KEY (id);


--
-- Name: multicast_group multicast_group_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.multicast_group
    ADD CONSTRAINT multicast_group_pkey PRIMARY KEY (id);


--
-- Name: network_server network_server_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.network_server
    ADD CONSTRAINT network_server_pkey PRIMARY KEY (id);


--
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: organization_user organization_user_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_pkey PRIMARY KEY (id);


--
-- Name: organization_user organization_user_user_id_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_user_id_organization_id_key UNIQUE (user_id, organization_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: service_profile service_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_pkey PRIMARY KEY (service_profile_id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: idx_api_key_application_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_api_key_application_id ON public.api_key USING btree (application_id);


--
-- Name: idx_api_key_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_api_key_organization_id ON public.api_key USING btree (organization_id);


--
-- Name: idx_application_name_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_application_name_trgm ON public.application USING gin (name public.gin_trgm_ops);


--
-- Name: idx_application_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_application_organization_id ON public.application USING btree (organization_id);


--
-- Name: idx_application_service_profile_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_application_service_profile_id ON public.application USING btree (service_profile_id);


--
-- Name: idx_device_application_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_application_id ON public.device USING btree (application_id);


--
-- Name: idx_device_dev_addr_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_dev_addr_trgm ON public.device USING gin (encode(dev_addr, 'hex'::text) public.gin_trgm_ops);


--
-- Name: idx_device_dev_eui_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_dev_eui_trgm ON public.device USING gin (encode(dev_eui, 'hex'::text) public.gin_trgm_ops);


--
-- Name: idx_device_device_profile_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_device_profile_id ON public.device USING btree (device_profile_id);


--
-- Name: idx_device_name_application_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE UNIQUE INDEX idx_device_name_application_id ON public.device USING btree (name, application_id);


--
-- Name: idx_device_name_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_name_trgm ON public.device USING gin (name public.gin_trgm_ops);


--
-- Name: idx_device_profile_network_server_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_profile_network_server_id ON public.device_profile USING btree (network_server_id);


--
-- Name: idx_device_profile_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_profile_organization_id ON public.device_profile USING btree (organization_id);


--
-- Name: idx_device_profile_tags; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_profile_tags ON public.device_profile USING gin (tags);


--
-- Name: idx_device_tags; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_device_tags ON public.device USING gin (tags);


--
-- Name: idx_gateway_gateway_profile_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_gateway_profile_id ON public.gateway USING btree (gateway_profile_id);


--
-- Name: idx_gateway_last_ping_sent_at; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_last_ping_sent_at ON public.gateway USING btree (last_ping_sent_at);


--
-- Name: idx_gateway_mac_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_mac_trgm ON public.gateway USING gin (encode(mac, 'hex'::text) public.gin_trgm_ops);


--
-- Name: idx_gateway_name_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE UNIQUE INDEX idx_gateway_name_organization_id ON public.gateway USING btree (name, organization_id);


--
-- Name: idx_gateway_name_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_name_trgm ON public.gateway USING gin (name public.gin_trgm_ops);


--
-- Name: idx_gateway_network_server_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_network_server_id ON public.gateway USING btree (network_server_id);


--
-- Name: idx_gateway_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_organization_id ON public.gateway USING btree (organization_id);


--
-- Name: idx_gateway_ping; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_ping ON public.gateway USING btree (ping);


--
-- Name: idx_gateway_ping_gateway_mac; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_ping_gateway_mac ON public.gateway_ping USING btree (gateway_mac);


--
-- Name: idx_gateway_ping_rx_gateway_mac; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_ping_rx_gateway_mac ON public.gateway_ping_rx USING btree (gateway_mac);


--
-- Name: idx_gateway_ping_rx_ping_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_ping_rx_ping_id ON public.gateway_ping_rx USING btree (ping_id);


--
-- Name: idx_gateway_profile_network_server_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_profile_network_server_id ON public.gateway_profile USING btree (network_server_id);


--
-- Name: idx_gateway_service_profile_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_service_profile_id ON public.gateway USING btree (service_profile_id);


--
-- Name: idx_gateway_tags; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_gateway_tags ON public.gateway USING gin (tags);


--
-- Name: idx_integration_application_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_integration_application_id ON public.integration USING btree (application_id);


--
-- Name: idx_integration_kind; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_integration_kind ON public.integration USING btree (kind);


--
-- Name: idx_multicast_group_application_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_multicast_group_application_id ON public.multicast_group USING btree (application_id);


--
-- Name: idx_multicast_group_name_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_multicast_group_name_trgm ON public.multicast_group USING gin (name public.gin_trgm_ops);


--
-- Name: idx_organization_name; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE UNIQUE INDEX idx_organization_name ON public.organization USING btree (name);


--
-- Name: idx_organization_name_trgm; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_organization_name_trgm ON public.organization USING gin (name public.gin_trgm_ops);


--
-- Name: idx_organization_user_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_organization_user_organization_id ON public.organization_user USING btree (organization_id);


--
-- Name: idx_organization_user_user_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_organization_user_user_id ON public.organization_user USING btree (user_id);


--
-- Name: idx_service_profile_network_server_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_service_profile_network_server_id ON public.service_profile USING btree (network_server_id);


--
-- Name: idx_service_profile_organization_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE INDEX idx_service_profile_organization_id ON public.service_profile USING btree (organization_id);


--
-- Name: idx_user_email; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE UNIQUE INDEX idx_user_email ON public."user" USING btree (email);


--
-- Name: idx_user_external_id; Type: INDEX; Schema: public; Owner: chirpstack_as
--

CREATE UNIQUE INDEX idx_user_external_id ON public."user" USING btree (external_id);


--
-- Name: api_key api_key_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.api_key
    ADD CONSTRAINT api_key_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: api_key api_key_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.api_key
    ADD CONSTRAINT api_key_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: application application_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: application application_service_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_service_profile_id_fkey FOREIGN KEY (service_profile_id) REFERENCES public.service_profile(service_profile_id);


--
-- Name: device device_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id);


--
-- Name: device device_device_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_device_profile_id_fkey FOREIGN KEY (device_profile_id) REFERENCES public.device_profile(device_profile_id);


--
-- Name: device_keys device_keys_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_keys
    ADD CONSTRAINT device_keys_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: device_multicast_group device_multicast_group_dev_eui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_multicast_group
    ADD CONSTRAINT device_multicast_group_dev_eui_fkey FOREIGN KEY (dev_eui) REFERENCES public.device(dev_eui) ON DELETE CASCADE;


--
-- Name: device_multicast_group device_multicast_group_multicast_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_multicast_group
    ADD CONSTRAINT device_multicast_group_multicast_group_id_fkey FOREIGN KEY (multicast_group_id) REFERENCES public.multicast_group(id) ON DELETE CASCADE;


--
-- Name: device_profile device_profile_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: device_profile device_profile_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.device_profile
    ADD CONSTRAINT device_profile_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);


--
-- Name: gateway gateway_gateway_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_gateway_profile_id_fkey FOREIGN KEY (gateway_profile_id) REFERENCES public.gateway_profile(gateway_profile_id);


--
-- Name: gateway gateway_last_ping_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_last_ping_id_fkey FOREIGN KEY (last_ping_id) REFERENCES public.gateway_ping(id) ON DELETE SET NULL;


--
-- Name: gateway gateway_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: gateway gateway_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: gateway_ping gateway_ping_gateway_mac_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping
    ADD CONSTRAINT gateway_ping_gateway_mac_fkey FOREIGN KEY (gateway_mac) REFERENCES public.gateway(mac) ON DELETE CASCADE;


--
-- Name: gateway_ping_rx gateway_ping_rx_gateway_mac_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_gateway_mac_fkey FOREIGN KEY (gateway_mac) REFERENCES public.gateway(mac) ON DELETE CASCADE;


--
-- Name: gateway_ping_rx gateway_ping_rx_ping_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_ping_rx
    ADD CONSTRAINT gateway_ping_rx_ping_id_fkey FOREIGN KEY (ping_id) REFERENCES public.gateway_ping(id) ON DELETE CASCADE;


--
-- Name: gateway_profile gateway_profile_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway_profile
    ADD CONSTRAINT gateway_profile_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: gateway gateway_service_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_service_profile_id_fkey FOREIGN KEY (service_profile_id) REFERENCES public.service_profile(service_profile_id);


--
-- Name: integration integration_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.integration
    ADD CONSTRAINT integration_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: multicast_group multicast_group_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.multicast_group
    ADD CONSTRAINT multicast_group_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- Name: organization_user organization_user_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) ON DELETE CASCADE;


--
-- Name: organization_user organization_user_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.organization_user
    ADD CONSTRAINT organization_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: service_profile service_profile_network_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_network_server_id_fkey FOREIGN KEY (network_server_id) REFERENCES public.network_server(id);


--
-- Name: service_profile service_profile_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chirpstack_as
--

ALTER TABLE ONLY public.service_profile
    ADD CONSTRAINT service_profile_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);


--
-- PostgreSQL database dump complete
--

