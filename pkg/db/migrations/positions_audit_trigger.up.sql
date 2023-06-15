CREATE OR REPLACE FUNCTION POSITIONS_AUDIT_TRIGGER(
) RETURNS TRIGGER AS 
	$$ BEGIN IF TG_OP = 'INSERT' THEN
	INSERT INTO
	    positions_audit (
	        id,
	        action,
	        title,
	        updated_at,
	        user_id,
	        position_id
	    )
	VALUES (
	        NEW.id,
	        'INSERT',
	        NEW.title,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'UPDATE' THEN
	INSERT INTO
	    positions_audit (
	        id,
	        action,
	        title,
	        updated_at,
	        user_id,
	        position_id
	    )
	VALUES (
	        NEW.id,
	        'UPDATE',
	        NEW.title,
	        NOW(),
	        NEW.user_id,
	        NEW.id
	    );
	ELSIF TG_OP = 'DELETE' THEN
	INSERT INTO
	    positions_audit (
	        id,
	        action,
	        title,
	        updated_at,
	        user_id,
	        position_id
	    )
	VALUES (
	        OLD.id,
	        'DELETE',
	        OLD.title,
	        NOW(),
	        OLD.user_id,
	        OLD.id
	    );
	END IF;
	RETURN 
NEW; 

END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER POSITIONS_AUDIT_TRIGGER 
	AFTER
	INSERT OR
	UPDATE OR
	DELETE ON positions FOR EACH ROW
	EXECUTE
	    FUNCTION positions_audit_trigger();
