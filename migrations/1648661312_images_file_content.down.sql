ALTER TABLE "images" DROP COLUMN "file_name", DROP COLUMN "content",
ADD COLUMN "message" character varying NOT NULL, ADD COLUMN "lsb_used" bigint NOT NULL, ADD COLUMN "channel" character varying NOT NULL
