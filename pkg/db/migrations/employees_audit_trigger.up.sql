CREATE OR REPLACE FUNCTION EMPLOYEES_AUDIT_TRIGGER(
) RETURNS TRIGGER AS 
	$$ BEGIN IF TG_OP = 'INSERT' THEN
	INSERT INTO
	    employees_audit (
	        id,
	        action,
	        first_name,
	        last_name,
	        email,
	        hire_date,
	        salary,
	        updated_at,
	        user_id,
	        employee_id
	    )
	VALUES (
	        NEW.id,
	        'INSERT',
	        NEW.first_name,
	        NEW.last_name,
	        NEW.email,
	        NEW.hire_date,
	        NEW.salary,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'UPDATE' THEN
	INSERT INTO
	    employees_audit (
	        id,
	        action,
	        first_name,
	        last_name,
	        email,
	        hire_date,
	        salary,
	        updated_at,
	        user_id,
	        employee_id
	    )
	VALUES (
	        NEW.id,
	        'UPDATE',
	        NEW.first_name,
	        NEW.last_name,
	        NEW.email,
	        NEW.hire_date,
	        NEW.salary,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'DELETE' THEN
	INSERT INTO
	    employees_audit (
	        id,
	        action,
	        first_name,
	        last_name,
	        email,
	        hire_date,
	        salary,
	        updated_at,
	        user_id,
	        employee_id
	    )
	VALUES (
	        OLD.id,
	        'DELETE',
	        OLD.first_name,
	        OLD.last_name,
	        OLD.email,
	        OLD.hire_date,
	        OLD.salary,
	        NOW(),
	        OLD.user_id,
	        OLD.id
	    );
	END IF;
	RETURN 
NEW; 

END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER EMPLOYEES_AUDIT_TRIGGER 
	AFTER
	INSERT OR
	UPDATE OR
	DELETE ON employees FOR EACH ROW
	EXECUTE
	    FUNCTION employees_audit_trigger();
