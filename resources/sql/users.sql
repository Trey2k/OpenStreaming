CREATE TABLE public.users (
    id bigint NOT NULL,
    "twitchID" character varying(80),
    "discordID" character varying(80),
    "refreshToken" character varying(80)
);

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);