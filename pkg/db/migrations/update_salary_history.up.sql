CREATE OR REPLACE FUNCTION UPDATE_SALARY_HISTORY() 
RETURNS TRIGGER AS 
	$$ BEGIN IF TG_OP = 'INSERT' THEN
	INSERT INTO
	    salary_history (
	        employee_id,
	        salary,
	        effective_date
	    )
	VALUES (
	        NEW.employee_id,
	        NEW.salary,
	        NEW.effective_date
	    );
	ELSIF TG_OP = 'UPDATE' THEN
	UPDATE salary_history
	SET end_date = NOW()
	WHERE
	    employee_id = NEW.employee_id
	    AND end_date IS NULL;
	INSERT INTO
	    salary_history (
	        employee_id,
	        salary,
	        effective_date
	    )
	VALUES (
	        NEW.employee_id,
	        NEW.salary,
	        NOW()
	    );
	END IF;
	RETURN 
NEW; 

END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER SALARY_HISTORY_TRIGGER 
	AFTER
	INSERT OR
	UPDATE
	    ON employee_salary_history FOR EACH ROW
	EXECUTE
	    FUNCTION update_salary_history();
