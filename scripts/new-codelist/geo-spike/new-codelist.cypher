CREATE CONSTRAINT ON (n:`_code_mcty16`) ASSERT n.value IS UNIQUE;

CREATE (node:`_code_list`:`_code_list_mcty16` { label:'Metropolitan Counties', edition:'December 2016', geography:true });

MERGE (node:`_code`:`_code_geography` { value:'E11000001' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000001' }) MERGE (node)-[:usedBy { label:"Greater Manchester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E11000002' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000002' }) MERGE (node)-[:usedBy { label:"Merseyside"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E11000003' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000003' }) MERGE (node)-[:usedBy { label:"South Yorkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E11000005' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000005' }) MERGE (node)-[:usedBy { label:"West Midlands"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E11000006' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000006' }) MERGE (node)-[:usedBy { label:"West Yorkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E11000007' });

MATCH (parent:`_code_list`:`_code_list_mcty16`),(node:`_code`:`_code_geography` { value:'E11000007' }) MERGE (node)-[:usedBy { label:"Tyne and Wear"}]->(parent);