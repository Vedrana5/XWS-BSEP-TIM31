INSERT INTO public.users(
	date_of_birth, email, first_name, is_confirmed, last_name, password, phone_number, recovery_email, role, username,is_using_fa)
	VALUES ('1999-06-12 00:00:00', 'admin123@gmail.com', 'Vedrana', true, 'Mijatovic', '$2a$10$q/8SyLsTjkr8yNvB8/L3jeNvQ7PsCqDS57cjGHld.1OdeOYMbAIFO', '066545545', 'vecaaa5@gmail.com',0, 'admin123',false);

INSERT INTO public.users(
	date_of_birth, email, first_name, is_confirmed, last_name, password, phone_number, recovery_email, role, username,is_using_fa)
	VALUES ('1999-06-12 00:00:00', 'vlasnik123@gmail.com', 'Vedrana', true, 'Mijatovic', '$2a$10$q/8SyLsTjkr8yNvB8/L3jeNvQ7PsCqDS57cjGHld.1OdeOYMbAIFO', '066545545', 'vecaa5@gmail.com',2, 'vlasnik123',false);

INSERT INTO public.users(
	date_of_birth, email, first_name, is_confirmed, last_name, password, phone_number, recovery_email, role, username,is_using_fa)
	VALUES ('1999-06-12 00:00:00', 'korisnik123@gmail.com', 'Vedrana', true, 'Mijatovic', '$2a$10$q/8SyLsTjkr8yNvB8/L3jeNvQ7PsCqDS57cjGHld.1OdeOYMbAIFO', '066545545', 'veca5@gmail.com',1, 'korisnik123',false);


INSERT INTO public.companies(
	name, website, email, phone_number, country_of_origin, founder, number_of_empl, owner_id,company_status)
	VALUES ('Glupost', 'Glupost jos veca','veca@gmail.com', '065756845','Srbija', 'Milan Ilic', '10', 2,0);

INSERT INTO public.companies(
	name, website, email, phone_number, country_of_origin, founder, number_of_empl, owner_id,company_status)
	VALUES ('Zahtjev', 'Zahtjev2','veca@gmail.com', '065756845','Srbija', 'Milan Ilic', '10', 2,2);

	INSERT INTO public.offers(position, description, date_created,due_date, company_id)
    	VALUES ('fdsfsdf', 'Glupost jos veca','2022-06-12 00:00:00', '2022-07-12 00:00:00',1);



 INSERT INTO permission(id, name,role)
 VALUES (2, 'CREATE_COMPANY_PERMISSION',2);

 INSERT INTO permission(id, name,role)
 VALUES (3, 'UPDATE_COMPANY_PERMISSION',2);

 INSERT INTO permission(id, name,role)
 VALUES (4, 'ACTIVATE_COMPANY_PERMISSION',0);


INSERT INTO permission(id, name,role)
VALUES (11, 'CREATE_SALARY_PERMISSION',1);

INSERT INTO permission(id, name,role)
VALUES (7, 'CREATE_JOB_AD_PERMISSION',1);

