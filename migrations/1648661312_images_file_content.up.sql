ALTER TABLE "images" ADD COLUMN "file_name" character varying NOT NULL, ADD COLUMN "content" character varying NOT NULL;
ALTER TABLE "images" DROP COLUMN "message", DROP COLUMN "lsb_used", DROP COLUMN "channel"
