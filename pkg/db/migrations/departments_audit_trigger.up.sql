CREATE OR REPLACE FUNCTION DEPARTMENTS_AUDIT_TRIGGER
() RETURNS TRIGGER AS 
	$$ BEGIN IF TG_OP = 'INSERT' THEN
	INSERT INTO
	    departments_audit (
	        id,
	        action,
	        name,
	        updated_at,
	        user_id,
	        department_id
	    )
	VALUES (
	        NEW.id,
	        'INSERT',
	        NEW.name,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'UPDATE' THEN
	INSERT INTO
	    departments_audit (
	        id,
	        action,
	        name,
	        updated_at,
	        user_id,
	        department_id
	    )
	VALUES (
	        NEW.id,
	        'UPDATE',
	        NEW.name,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'DELETE' THEN
	INSERT INTO
	    departments_audit (
	        id,
	        action,
	        name,
	        updated_at,
	        user_id,
	        department_id
	    )
	VALUES (
	        OLD.id,
	        'DELETE',
	        OLD.name,
	        NOW(),
	        OLD.user_id,
	        OLD.id
	    );
	END IF;
	RETURN 
NEW; 

END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER DEPARTMENTS_AUDIT_TRIGGER 
	AFTER
	INSERT OR
	UPDATE OR
	DELETE
	    ON departments FOR EACH ROW
	EXECUTE
	    FUNCTION departments_audit_trigger();
