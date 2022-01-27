CREATE TABLE public.overlays (
    id bigint NOT NULL,
    "userID" bigint NOT NULL,
    "key" character varying(80) NOT NULL
);

CREATE SEQUENCE public.overlays_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.overlays_id_seq OWNED BY public.overlays.id;
ALTER TABLE ONLY public.overlays ALTER COLUMN id SET DEFAULT nextval('public.overlays_id_seq'::regclass);