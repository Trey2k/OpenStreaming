CREATE TABLE public."overlayModules" (
    id bigint NOT NULL,
    "overlayID" bigint NOT NULL,
    "type" bigint NOT NULL,
    "top" bigint NOT NULL,
    "left" bigint NOT NULL,
    "width" bigint NOT NULL,
    "height" bigint NOT NULL
);

CREATE SEQUENCE public."overlayModules_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public."overlayModules_id_seq" OWNED BY public."overlayModules".id;
ALTER TABLE ONLY public."overlayModules" ALTER COLUMN id SET DEFAULT nextval('public."overlayModules_id_seq"'::regclass);