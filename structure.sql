CREATE SEQUENCE soccer_team_id_seq
    INCREMENT 1
    START 10
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE soccer_team
(
    id bigint NOT NULL DEFAULT nextval('soccer_team_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    coach_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT soccer_team_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

CREATE SEQUENCE soccer_player_id_seq
    INCREMENT 1
    START 10
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE soccer_player
(
    id bigint NOT NULL DEFAULT nextval('soccer_player_id_seq'::regclass),
    team_id bigint NOT NULL,
    fullname character varying(100) COLLATE pg_catalog."default" NOT NULL,
    dob date NOT NULL,
    age bigint NOT NULL,
    nationality character varying(50) COLLATE pg_catalog."default" NOT NULL,
    "position" character varying(10) COLLATE pg_catalog."default" NOT NULL,
    shirt_number character varying(5) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT soccer_player_pk PRIMARY KEY (id),
    CONSTRAINT fk_soccer_player__team FOREIGN KEY (team_id)
        REFERENCES soccer_team (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

