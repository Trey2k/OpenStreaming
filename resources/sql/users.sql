CREATE TABLE public.viewers (
    id bigint NOT NULL,
    "twitchID" character varying(80),
    "discordID" character varying(80),
    "refreshToken" character varying(80),
    "expires"
    "hasFollowed" boolean DEFAULT false NOT NULL
);

CREATE SEQUENCE public.viewers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.viewers_id_seq OWNED BY public.viewers.id;
ALTER TABLE ONLY public.viewers ALTER COLUMN id SET DEFAULT nextval('public.viewers_id_seq'::regclass);