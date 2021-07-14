-- Table: public.aboutme

-- DROP TABLE public.aboutme;

CREATE TABLE public.aboutme
(
    id integer NOT NULL DEFAULT nextval('aboutme_id_seq'::regclass),
    title text COLLATE pg_catalog."default",
    info text COLLATE pg_catalog."default",
    CONSTRAINT aboutme_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.aboutme
    OWNER to postgres;