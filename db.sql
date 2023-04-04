
-- CREATE TABLE "product_images" (
-- 	"id" serial NOT NULL,
-- 	"product_id" integer NOT NULL,
-- 	"img" varchar(255) NOT NULL,
-- 	CONSTRAINT "product_images_pk" PRIMARY KEY ("id")
-- ) WITH (
--   OIDS=FALSE
-- );


-- ALTER TABLE "product_images" ADD CONSTRAINT "product_images_fk0" FOREIGN KEY ("product_id") REFERENCES "products"("id");


INSERT INTO public.products (title, "desc", price, img, "type") VALUES('Test Product 1', 'Test description', 10, NULL, 'mobile');
INSERT INTO public.products (title, "desc", price, img, "type") VALUES('Test Product 2', 'Test description', 10, NULL, 'mobile');
INSERT INTO public.products (title, "desc", price, img, "type") VALUES('Test Product 3', 'Test description', 10, NULL, 'tv');
INSERT INTO public.products (title, "desc", price, img, "type") VALUES('Test Product 4', 'Test description', 10, NULL, 'pc');
