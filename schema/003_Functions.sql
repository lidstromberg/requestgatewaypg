/*********************************************************************
-- FUNCTION: public.set_gateway(character varying,character varying)
-- DROP FUNCTION public.set_gateway(character varying,character varying);
*********************************************************************/

CREATE OR REPLACE FUNCTION public.set_gateway(
	in_appscope character varying(510),
    in_remoteaddress character varying(510))
    RETURNS void
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
AS $BODY$
/*********************************************************************
Name: set_gateway
Auth: DF
Date: 11.04.2019

Notes:
    sets a gateway address and application scope

*********************************************************************/
DECLARE 
    l_gatewayid bigint;
BEGIN
    --check if the record exists
    select count(1)
    into l_gatewayid
    from public.gateway g1
    where g1.appscope=in_appscope
    and g1.remoteaddress=in_remoteaddress;

    --exit if it's already present
    if l_gatewayid > 0 then
        return;
    end if;

    --otherwise insert it
    insert into public.gateway
    (
        appscope,
        remoteaddress
    )
    values
    (
        in_appscope,
        in_remoteaddress
    );
END

$BODY$;

ALTER FUNCTION public.set_gateway(character varying,character varying) OWNER TO postgres;
GRANT ALL ON FUNCTION public.set_gateway(character varying,character varying) to gatewayuser;


CREATE OR REPLACE FUNCTION public.get_gateway(
	in_appscope character varying(510),
    in_remoteaddress character varying(510))
    RETURNS boolean
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
AS $BODY$

/*********************************************************************
Name: get_gateway
Auth: DF
Date: 11.04.2019

Notes:
    Returns boolean indicator showing true if the address is cleared
    for the address

*********************************************************************/
DECLARE 
    l_gatewayid bigint;
BEGIN
    --check if the record exists
    select count(1)
    into l_gatewayid
    from public.gateway g1
    where g1.appscope=in_appscope
    and g1.remoteaddress=in_remoteaddress;

    --return true if it's present
    if l_gatewayid > 0 then
        return true;
    end if;

    --return false if it isn't present
	return false;
END

$BODY$;

ALTER FUNCTION public.get_gateway(character varying,character varying) OWNER TO postgres;
GRANT ALL ON FUNCTION public.get_gateway(character varying,character varying) to gatewayuser;


-- FUNCTION: public.delete_gateway(character varying)
-- DROP FUNCTION public.delete_gateway(character varying);

CREATE OR REPLACE FUNCTION public.delete_gateway(
	in_appscope character varying(510),
    in_remoteaddress character varying(510))
    RETURNS void
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
AS $BODY$
/*********************************************************************
Name: delete_gateway
Auth: DF
Date: 11.04.2019

Notes:
    Removes a gateway address

*********************************************************************/
DECLARE 
    l_gatewayid bigint;
BEGIN
    --check if the record exists
    select count(1)
    into l_gatewayid
    from public.gateway g1
    where g1.appscope=in_appscope
    and g1.remoteaddress=in_remoteaddress;

    --delete if it's present
    if l_gatewayid > 0 then
        delete from public.gateway g1
        where g1.appscope=in_appscope
        and g1.remoteaddress=in_remoteaddress;
    end if;

	return;
END

$BODY$;

ALTER FUNCTION public.delete_gateway(character varying,character varying) OWNER TO postgres;
GRANT ALL ON FUNCTION public.delete_gateway(character varying,character varying) to gatewayuser;


select public.set_gateway('testapp','0.0.0.1');
select public.get_gateway('testapp','0.0.0.1');
select public.delete_gateway('testapp','0.0.0.1');