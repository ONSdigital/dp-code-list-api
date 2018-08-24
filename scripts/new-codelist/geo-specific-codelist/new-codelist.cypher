CREATE CONSTRAINT ON (n:`_code_geography`) ASSERT n.value IS UNIQUE;

CREATE (node:`_code_list`:`_code_list_local-authority`:`_geography` { label:'Local Authority Districts', edition:'2016', geography_list_ID:'lad16' });

MERGE (node:`_code`:`_code_geography` { value:'E09000001' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000001' }) MERGE (node)-[:usedBy { label:"City of London"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000002' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000002' }) MERGE (node)-[:usedBy { label:"Barking and Dagenham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000003' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000003' }) MERGE (node)-[:usedBy { label:"Barnet"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000004' }) MERGE (node)-[:usedBy { label:"Bexley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000005' }) MERGE (node)-[:usedBy { label:"Brent"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000006' }) MERGE (node)-[:usedBy { label:"Bromley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000007' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000007' }) MERGE (node)-[:usedBy { label:"Camden"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000008' }) MERGE (node)-[:usedBy { label:"Croydon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000009' }) MERGE (node)-[:usedBy { label:"Ealing"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000010' }) MERGE (node)-[:usedBy { label:"Enfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000024' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000024' }) MERGE (node)-[:usedBy { label:"Sunderland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000011' }) MERGE (node)-[:usedBy { label:"Greenwich"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000012' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000012' }) MERGE (node)-[:usedBy { label:"Hackney"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000025' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000025' }) MERGE (node)-[:usedBy { label:"Birmingham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000243' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000243' }) MERGE (node)-[:usedBy { label:"Stevenage"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000026' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000026' }) MERGE (node)-[:usedBy { label:"Coventry"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000013' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000013' }) MERGE (node)-[:usedBy { label:"Hammersmith and Fulham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000014' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000014' }) MERGE (node)-[:usedBy { label:"Haringey"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000027' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000027' }) MERGE (node)-[:usedBy { label:"Dudley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000102' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000102' }) MERGE (node)-[:usedBy { label:"Three Rivers"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000028' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000028' }) MERGE (node)-[:usedBy { label:"Sandwell"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000015' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000015' }) MERGE (node)-[:usedBy { label:"Harrow"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000103' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000103' }) MERGE (node)-[:usedBy { label:"Watford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000034' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000034' }) MERGE (node)-[:usedBy { label:"Thurrock"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000029' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000029' }) MERGE (node)-[:usedBy { label:"Solihull"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000016' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000016' }) MERGE (node)-[:usedBy { label:"Havering"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000035' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000035' }) MERGE (node)-[:usedBy { label:"Medway"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000241' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000241' }) MERGE (node)-[:usedBy { label:"Welwyn Hatfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000017' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000017' }) MERGE (node)-[:usedBy { label:"Hillingdon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000030' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000030' }) MERGE (node)-[:usedBy { label:"Walsall"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000036' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000036' }) MERGE (node)-[:usedBy { label:"Bracknell Forest"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000038' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000038' }) MERGE (node)-[:usedBy { label:"North East Derbyshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000039' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000039' }) MERGE (node)-[:usedBy { label:"South Derbyshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000037' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000037' }) MERGE (node)-[:usedBy { label:"West Berkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000031' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000031' }) MERGE (node)-[:usedBy { label:"Wolverhampton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000105' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000105' }) MERGE (node)-[:usedBy { label:"Ashford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000018' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000018' }) MERGE (node)-[:usedBy { label:"Hounslow"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000040' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000040' }) MERGE (node)-[:usedBy { label:"East Devon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000106' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000106' }) MERGE (node)-[:usedBy { label:"Canterbury"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000019' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000019' }) MERGE (node)-[:usedBy { label:"Islington"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000032' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000032' }) MERGE (node)-[:usedBy { label:"Bradford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000038' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000038' }) MERGE (node)-[:usedBy { label:"Reading"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000041' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000041' }) MERGE (node)-[:usedBy { label:"Exeter"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000107' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000107' }) MERGE (node)-[:usedBy { label:"Dartford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000033' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000033' }) MERGE (node)-[:usedBy { label:"Calderdale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000020' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000020' }) MERGE (node)-[:usedBy { label:"Kensington and Chelsea"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000039' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000039' }) MERGE (node)-[:usedBy { label:"Slough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000042' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000042' }) MERGE (node)-[:usedBy { label:"Mid Devon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000108' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000108' }) MERGE (node)-[:usedBy { label:"Dover"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000034' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000034' }) MERGE (node)-[:usedBy { label:"Kirklees"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000021' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000021' }) MERGE (node)-[:usedBy { label:"Kingston upon Thames"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000040' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000040' }) MERGE (node)-[:usedBy { label:"Windsor and Maidenhead"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000043' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000043' }) MERGE (node)-[:usedBy { label:"North Devon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000109' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000109' }) MERGE (node)-[:usedBy { label:"Gravesham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000035' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000035' }) MERGE (node)-[:usedBy { label:"Leeds"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000022' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000022' }) MERGE (node)-[:usedBy { label:"Lambeth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000041' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000041' }) MERGE (node)-[:usedBy { label:"Wokingham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000044' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000044' }) MERGE (node)-[:usedBy { label:"South Hams"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000036' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000036' }) MERGE (node)-[:usedBy { label:"Wakefield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000110' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000110' }) MERGE (node)-[:usedBy { label:"Maidstone"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000023' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000023' }) MERGE (node)-[:usedBy { label:"Lewisham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000045' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000045' }) MERGE (node)-[:usedBy { label:"Teignbridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000111' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000111' }) MERGE (node)-[:usedBy { label:"Sevenoaks"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000024' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000024' }) MERGE (node)-[:usedBy { label:"Merton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000001' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000001' }) MERGE (node)-[:usedBy { label:"Hartlepool"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000046' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000046' }) MERGE (node)-[:usedBy { label:"Torridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000042' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000042' }) MERGE (node)-[:usedBy { label:"Milton Keynes"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000112' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000112' }) MERGE (node)-[:usedBy { label:"Shepway"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000047' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000047' }) MERGE (node)-[:usedBy { label:"West Devon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000113' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000113' }) MERGE (node)-[:usedBy { label:"Swale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000224' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000224' }) MERGE (node)-[:usedBy { label:"Arun"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000043' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000043' }) MERGE (node)-[:usedBy { label:"Brighton and Hove"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000002' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000002' }) MERGE (node)-[:usedBy { label:"Middlesbrough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000025' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000025' }) MERGE (node)-[:usedBy { label:"Newham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000225' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000225' }) MERGE (node)-[:usedBy { label:"Chichester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000114' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000114' }) MERGE (node)-[:usedBy { label:"Thanet"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000048' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000048' }) MERGE (node)-[:usedBy { label:"Christchurch"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000044' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000044' }) MERGE (node)-[:usedBy { label:"Portsmouth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000226' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000226' }) MERGE (node)-[:usedBy { label:"Crawley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000049' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000049' }) MERGE (node)-[:usedBy { label:"East Dorset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000115' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000115' }) MERGE (node)-[:usedBy { label:"Tonbridge and Malling"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000163' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000163' }) MERGE (node)-[:usedBy { label:"Craven"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000045' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000045' }) MERGE (node)-[:usedBy { label:"Southampton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000026' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000026' }) MERGE (node)-[:usedBy { label:"Redbridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000003' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000003' }) MERGE (node)-[:usedBy { label:"Redcar and Cleveland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000050' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000050' }) MERGE (node)-[:usedBy { label:"North Dorset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000116' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000116' }) MERGE (node)-[:usedBy { label:"Tunbridge Wells"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000227' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000227' }) MERGE (node)-[:usedBy { label:"Horsham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000164' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000164' }) MERGE (node)-[:usedBy { label:"Hambleton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000004' }) MERGE (node)-[:usedBy { label:"Stockton-on-Tees"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000046' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000046' }) MERGE (node)-[:usedBy { label:"Isle of Wight"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000051' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000051' }) MERGE (node)-[:usedBy { label:"Purbeck"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000027' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000027' }) MERGE (node)-[:usedBy { label:"Richmond upon Thames"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000117' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000117' }) MERGE (node)-[:usedBy { label:"Burnley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000228' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000228' }) MERGE (node)-[:usedBy { label:"Mid Sussex"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000118' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000118' }) MERGE (node)-[:usedBy { label:"Chorley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000001' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000001' }) MERGE (node)-[:usedBy { label:"Isle of Anglesey"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000028' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000028' }) MERGE (node)-[:usedBy { label:"Southwark"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000052' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000052' }) MERGE (node)-[:usedBy { label:"West Dorset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000005' }) MERGE (node)-[:usedBy { label:"Darlington"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000165' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000165' }) MERGE (node)-[:usedBy { label:"Harrogate"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000229' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000229' }) MERGE (node)-[:usedBy { label:"Worthing"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000029' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000029' }) MERGE (node)-[:usedBy { label:"Sutton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000119' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000119' }) MERGE (node)-[:usedBy { label:"Fylde"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000047' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000047' }) MERGE (node)-[:usedBy { label:"County Durham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000053' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000053' }) MERGE (node)-[:usedBy { label:"Weymouth and Portland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000002' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000002' }) MERGE (node)-[:usedBy { label:"Gwynedd"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000166' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000166' }) MERGE (node)-[:usedBy { label:"Richmondshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000234' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000234' }) MERGE (node)-[:usedBy { label:"Bromsgrove"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000003' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000003' }) MERGE (node)-[:usedBy { label:"Conwy"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000120' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000120' }) MERGE (node)-[:usedBy { label:"Hyndburn"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000167' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000167' }) MERGE (node)-[:usedBy { label:"Ryedale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000061' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000061' }) MERGE (node)-[:usedBy { label:"Eastbourne"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000057' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000057' }) MERGE (node)-[:usedBy { label:"Northumberland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000030' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000030' }) MERGE (node)-[:usedBy { label:"Tower Hamlets"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000235' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000235' }) MERGE (node)-[:usedBy { label:"Malvern Hills"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000004' }) MERGE (node)-[:usedBy { label:"Denbighshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000121' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000121' }) MERGE (node)-[:usedBy { label:"Lancaster"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000062' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000062' }) MERGE (node)-[:usedBy { label:"Hastings"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000031' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000031' }) MERGE (node)-[:usedBy { label:"Waltham Forest"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000168' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000168' }) MERGE (node)-[:usedBy { label:"Scarborough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000049' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000049' }) MERGE (node)-[:usedBy { label:"Cheshire East"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000236' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000236' }) MERGE (node)-[:usedBy { label:"Redditch"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000005' }) MERGE (node)-[:usedBy { label:"Flintshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000122' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000122' }) MERGE (node)-[:usedBy { label:"Pendle"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000063' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000063' }) MERGE (node)-[:usedBy { label:"Lewes"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000032' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000032' }) MERGE (node)-[:usedBy { label:"Wandsworth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000169' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000169' }) MERGE (node)-[:usedBy { label:"Selby"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000006' }) MERGE (node)-[:usedBy { label:"Halton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000237' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000237' }) MERGE (node)-[:usedBy { label:"Worcester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000006' }) MERGE (node)-[:usedBy { label:"Wrexham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000064' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000064' }) MERGE (node)-[:usedBy { label:"Rother"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E09000033' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E09000033' }) MERGE (node)-[:usedBy { label:"Westminster"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000123' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000123' }) MERGE (node)-[:usedBy { label:"Preston"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000170' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000170' }) MERGE (node)-[:usedBy { label:"Ashfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000007' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000007' }) MERGE (node)-[:usedBy { label:"Warrington"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000023' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000023' }) MERGE (node)-[:usedBy { label:"Powys"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000238' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000238' }) MERGE (node)-[:usedBy { label:"Wychavon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000065' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000065' }) MERGE (node)-[:usedBy { label:"Wealden"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000124' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000124' }) MERGE (node)-[:usedBy { label:"Ribble Valley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000050' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000050' }) MERGE (node)-[:usedBy { label:"Cheshire West and Chester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000001' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000001' }) MERGE (node)-[:usedBy { label:"Bolton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000171' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000171' }) MERGE (node)-[:usedBy { label:"Bassetlaw"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000008' }) MERGE (node)-[:usedBy { label:"Ceredigion"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000125' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000125' }) MERGE (node)-[:usedBy { label:"Rossendale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000239' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000239' }) MERGE (node)-[:usedBy { label:"Wyre Forest"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000008' }) MERGE (node)-[:usedBy { label:"Blackburn with Darwen"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000066' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000066' }) MERGE (node)-[:usedBy { label:"Basildon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000002' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000002' }) MERGE (node)-[:usedBy { label:"Bury"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000172' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000172' }) MERGE (node)-[:usedBy { label:"Broxtowe"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000009' }) MERGE (node)-[:usedBy { label:"Pembrokeshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000005' }) MERGE (node)-[:usedBy { label:"Clackmannanshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000009' }) MERGE (node)-[:usedBy { label:"Blackpool"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000067' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000067' }) MERGE (node)-[:usedBy { label:"Braintree"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000003' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000003' }) MERGE (node)-[:usedBy { label:"Manchester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000126' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000126' }) MERGE (node)-[:usedBy { label:"South Ribble"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000173' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000173' }) MERGE (node)-[:usedBy { label:"Gedling"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000010' }) MERGE (node)-[:usedBy { label:"Carmarthenshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000006' }) MERGE (node)-[:usedBy { label:"Dumfries and Galloway"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000068' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000068' }) MERGE (node)-[:usedBy { label:"Brentwood"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000010' }) MERGE (node)-[:usedBy { label:"Kingston upon Hull, City of"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000174' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000174' }) MERGE (node)-[:usedBy { label:"Mansfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000127' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000127' }) MERGE (node)-[:usedBy { label:"West Lancashire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000004' }) MERGE (node)-[:usedBy { label:"Oldham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000011' }) MERGE (node)-[:usedBy { label:"Swansea"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000008' }) MERGE (node)-[:usedBy { label:"East Ayrshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000012' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000012' }) MERGE (node)-[:usedBy { label:"Neath Port Talbot"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000069' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000069' }) MERGE (node)-[:usedBy { label:"Castle Point"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000011' }) MERGE (node)-[:usedBy { label:"East Riding of Yorkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000128' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000128' }) MERGE (node)-[:usedBy { label:"Wyre"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000175' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000175' }) MERGE (node)-[:usedBy { label:"Newark and Sherwood"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000005' }) MERGE (node)-[:usedBy { label:"Rochdale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000010' }) MERGE (node)-[:usedBy { label:"East Lothian"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000012' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000012' }) MERGE (node)-[:usedBy { label:"North East Lincolnshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000129' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000129' }) MERGE (node)-[:usedBy { label:"Blaby"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000070' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000070' }) MERGE (node)-[:usedBy { label:"Chelmsford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000176' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000176' }) MERGE (node)-[:usedBy { label:"Rushcliffe"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000006' }) MERGE (node)-[:usedBy { label:"Salford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000013' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000013' }) MERGE (node)-[:usedBy { label:"Bridgend"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000013' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000013' }) MERGE (node)-[:usedBy { label:"North Lincolnshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000011' }) MERGE (node)-[:usedBy { label:"East Renfrewshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000130' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000130' }) MERGE (node)-[:usedBy { label:"Charnwood"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000071' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000071' }) MERGE (node)-[:usedBy { label:"Colchester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000177' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000177' }) MERGE (node)-[:usedBy { label:"Cherwell"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000007' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000007' }) MERGE (node)-[:usedBy { label:"Stockport"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000014' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000014' }) MERGE (node)-[:usedBy { label:"Vale of Glamorgan"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000013' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000013' }) MERGE (node)-[:usedBy { label:"Na h-Eileanan Siar"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000178' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000178' }) MERGE (node)-[:usedBy { label:"Oxford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000131' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000131' }) MERGE (node)-[:usedBy { label:"Harborough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000016' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000016' }) MERGE (node)-[:usedBy { label:"Rhondda Cynon Taf"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000072' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000072' }) MERGE (node)-[:usedBy { label:"Epping Forest"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000008' }) MERGE (node)-[:usedBy { label:"Tameside"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000014' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000014' }) MERGE (node)-[:usedBy { label:"York"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000014' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000014' }) MERGE (node)-[:usedBy { label:"Falkirk"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000024' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000024' }) MERGE (node)-[:usedBy { label:"Merthyr Tydfil"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000073' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000073' }) MERGE (node)-[:usedBy { label:"Harlow"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000009' }) MERGE (node)-[:usedBy { label:"Trafford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000132' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000132' }) MERGE (node)-[:usedBy { label:"Hinckley and Bosworth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000015' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000015' }) MERGE (node)-[:usedBy { label:"Derby"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000179' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000179' }) MERGE (node)-[:usedBy { label:"South Oxfordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000015' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000015' }) MERGE (node)-[:usedBy { label:"Fife"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000018' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000018' }) MERGE (node)-[:usedBy { label:"Caerphilly"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000074' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000074' }) MERGE (node)-[:usedBy { label:"Maldon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000133' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000133' }) MERGE (node)-[:usedBy { label:"Melton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000016' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000016' }) MERGE (node)-[:usedBy { label:"Leicester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000180' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000180' }) MERGE (node)-[:usedBy { label:"Vale of White Horse"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000010' }) MERGE (node)-[:usedBy { label:"Wigan"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000017' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000017' }) MERGE (node)-[:usedBy { label:"Highland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000017' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000017' }) MERGE (node)-[:usedBy { label:"Rutland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000075' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000075' }) MERGE (node)-[:usedBy { label:"Rochford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000181' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000181' }) MERGE (node)-[:usedBy { label:"West Oxfordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000134' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000134' }) MERGE (node)-[:usedBy { label:"North West Leicestershire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000011' }) MERGE (node)-[:usedBy { label:"Knowsley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000019' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000019' }) MERGE (node)-[:usedBy { label:"Blaenau Gwent"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000018' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000018' }) MERGE (node)-[:usedBy { label:"Inverclyde"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000018' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000018' }) MERGE (node)-[:usedBy { label:"Nottingham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000076' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000076' }) MERGE (node)-[:usedBy { label:"Tendring"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000012' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000012' }) MERGE (node)-[:usedBy { label:"Liverpool"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000020' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000020' }) MERGE (node)-[:usedBy { label:"Torfaen"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000135' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000135' }) MERGE (node)-[:usedBy { label:"Oadby and Wigston"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000187' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000187' }) MERGE (node)-[:usedBy { label:"Mendip"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000019' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000019' }) MERGE (node)-[:usedBy { label:"Midlothian"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000019' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000019' }) MERGE (node)-[:usedBy { label:"Herefordshire, County of"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000077' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000077' }) MERGE (node)-[:usedBy { label:"Uttlesford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000188' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000188' }) MERGE (node)-[:usedBy { label:"Sedgemoor"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000136' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000136' }) MERGE (node)-[:usedBy { label:"Boston"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000013' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000013' }) MERGE (node)-[:usedBy { label:"St. Helens"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000021' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000021' }) MERGE (node)-[:usedBy { label:"Monmouthshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000020' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000020' }) MERGE (node)-[:usedBy { label:"Moray"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000020' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000020' }) MERGE (node)-[:usedBy { label:"Telford and Wrekin"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000078' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000078' }) MERGE (node)-[:usedBy { label:"Cheltenham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000137' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000137' }) MERGE (node)-[:usedBy { label:"East Lindsey"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000014' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000014' }) MERGE (node)-[:usedBy { label:"Sefton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000189' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000189' }) MERGE (node)-[:usedBy { label:"South Somerset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000022' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000022' }) MERGE (node)-[:usedBy { label:"Newport"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000021' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000021' }) MERGE (node)-[:usedBy { label:"North Ayrshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000051' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000051' }) MERGE (node)-[:usedBy { label:"Shropshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000138' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000138' }) MERGE (node)-[:usedBy { label:"Lincoln"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000015' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000015' }) MERGE (node)-[:usedBy { label:"Wirral"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000079' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000079' }) MERGE (node)-[:usedBy { label:"Cotswold"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'W06000015' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'W06000015' }) MERGE (node)-[:usedBy { label:"Cardiff"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000190' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000190' }) MERGE (node)-[:usedBy { label:"Taunton Deane"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000023' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000023' }) MERGE (node)-[:usedBy { label:"Orkney Islands"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000021' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000021' }) MERGE (node)-[:usedBy { label:"Stoke-on-Trent"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000080' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000080' }) MERGE (node)-[:usedBy { label:"Forest of Dean"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000016' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000016' }) MERGE (node)-[:usedBy { label:"Barnsley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000139' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000139' }) MERGE (node)-[:usedBy { label:"North Kesteven"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000191' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000191' }) MERGE (node)-[:usedBy { label:"West Somerset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000004' }) MERGE (node)-[:usedBy { label:"Aylesbury Vale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000024' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000024' }) MERGE (node)-[:usedBy { label:"Perth and Kinross"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000022' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000022' }) MERGE (node)-[:usedBy { label:"Bath and North East Somerset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000081' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000081' }) MERGE (node)-[:usedBy { label:"Gloucester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000017' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000017' }) MERGE (node)-[:usedBy { label:"Doncaster"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000140' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000140' }) MERGE (node)-[:usedBy { label:"South Holland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000192' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000192' }) MERGE (node)-[:usedBy { label:"Cannock Chase"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000005' }) MERGE (node)-[:usedBy { label:"Chiltern"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000026' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000026' }) MERGE (node)-[:usedBy { label:"Scottish Borders"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000082' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000082' }) MERGE (node)-[:usedBy { label:"Stroud"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000023' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000023' }) MERGE (node)-[:usedBy { label:"Bristol, City of"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000018' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000018' }) MERGE (node)-[:usedBy { label:"Rotherham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000193' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000193' }) MERGE (node)-[:usedBy { label:"East Staffordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000141' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000141' }) MERGE (node)-[:usedBy { label:"South Kesteven"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000006' }) MERGE (node)-[:usedBy { label:"South Bucks"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000027' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000027' }) MERGE (node)-[:usedBy { label:"Shetland Islands"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000142' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000142' }) MERGE (node)-[:usedBy { label:"West Lindsey"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000007' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000007' }) MERGE (node)-[:usedBy { label:"Wycombe"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000019' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000019' }) MERGE (node)-[:usedBy { label:"Sheffield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000083' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000083' }) MERGE (node)-[:usedBy { label:"Tewkesbury"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000194' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000194' }) MERGE (node)-[:usedBy { label:"Lichfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000024' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000024' }) MERGE (node)-[:usedBy { label:"North Somerset"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000028' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000028' }) MERGE (node)-[:usedBy { label:"South Ayrshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000143' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000143' }) MERGE (node)-[:usedBy { label:"Breckland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000008' }) MERGE (node)-[:usedBy { label:"Cambridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000084' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000084' }) MERGE (node)-[:usedBy { label:"Basingstoke and Deane"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000037' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000037' }) MERGE (node)-[:usedBy { label:"Gateshead"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000195' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000195' }) MERGE (node)-[:usedBy { label:"Newcastle-under-Lyme"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000025' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000025' }) MERGE (node)-[:usedBy { label:"South Gloucestershire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000029' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000029' }) MERGE (node)-[:usedBy { label:"South Lanarkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000144' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000144' }) MERGE (node)-[:usedBy { label:"Broadland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000009' }) MERGE (node)-[:usedBy { label:"East Cambridgeshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000052' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000052' }) MERGE (node)-[:usedBy { label:"Cornwall"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000085' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000085' }) MERGE (node)-[:usedBy { label:"East Hampshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000196' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000196' }) MERGE (node)-[:usedBy { label:"South Staffordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000021' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000021' }) MERGE (node)-[:usedBy { label:"Newcastle upon Tyne"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000030' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000030' }) MERGE (node)-[:usedBy { label:"Stirling"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000197' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000197' }) MERGE (node)-[:usedBy { label:"Stafford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000086' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000086' }) MERGE (node)-[:usedBy { label:"Eastleigh"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000022' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000022' }) MERGE (node)-[:usedBy { label:"North Tyneside"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000145' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000145' }) MERGE (node)-[:usedBy { label:"Great Yarmouth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000010' }) MERGE (node)-[:usedBy { label:"Fenland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000053' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000053' }) MERGE (node)-[:usedBy { label:"Isles of Scilly"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000033' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000033' }) MERGE (node)-[:usedBy { label:"Aberdeen City"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000198' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000198' }) MERGE (node)-[:usedBy { label:"Staffordshire Moorlands"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000087' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000087' }) MERGE (node)-[:usedBy { label:"Fareham"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E08000023' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E08000023' }) MERGE (node)-[:usedBy { label:"South Tyneside"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000146' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000146' }) MERGE (node)-[:usedBy { label:"King's Lynn and West Norfolk"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000011' }) MERGE (node)-[:usedBy { label:"Huntingdonshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000026' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000026' }) MERGE (node)-[:usedBy { label:"Plymouth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000034' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000034' }) MERGE (node)-[:usedBy { label:"Aberdeenshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000199' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000199' }) MERGE (node)-[:usedBy { label:"Tamworth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000012' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000012' }) MERGE (node)-[:usedBy { label:"South Cambridgeshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000147' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000147' }) MERGE (node)-[:usedBy { label:"North Norfolk"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000088' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000088' }) MERGE (node)-[:usedBy { label:"Gosport"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000027' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000027' }) MERGE (node)-[:usedBy { label:"Torbay"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000035' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000035' }) MERGE (node)-[:usedBy { label:"Argyll and Bute"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000200' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000200' }) MERGE (node)-[:usedBy { label:"Babergh"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000148' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000148' }) MERGE (node)-[:usedBy { label:"Norwich"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000026' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000026' }) MERGE (node)-[:usedBy { label:"Allerdale"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000089' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000089' }) MERGE (node)-[:usedBy { label:"Hart"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000028' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000028' }) MERGE (node)-[:usedBy { label:"Bournemouth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000036' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000036' }) MERGE (node)-[:usedBy { label:"City of Edinburgh"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000201' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000201' }) MERGE (node)-[:usedBy { label:"Forest Heath"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000149' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000149' }) MERGE (node)-[:usedBy { label:"South Norfolk"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000029' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000029' }) MERGE (node)-[:usedBy { label:"Poole"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000090' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000090' }) MERGE (node)-[:usedBy { label:"Havant"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000027' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000027' }) MERGE (node)-[:usedBy { label:"Barrow-in-Furness"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000038' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000038' }) MERGE (node)-[:usedBy { label:"Renfrewshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000150' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000150' }) MERGE (node)-[:usedBy { label:"Corby"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000202' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000202' }) MERGE (node)-[:usedBy { label:"Ipswich"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000028' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000028' }) MERGE (node)-[:usedBy { label:"Carlisle"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000091' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000091' }) MERGE (node)-[:usedBy { label:"New Forest"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000030' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000030' }) MERGE (node)-[:usedBy { label:"Swindon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000039' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000039' }) MERGE (node)-[:usedBy { label:"West Dunbartonshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000151' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000151' }) MERGE (node)-[:usedBy { label:"Daventry"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000203' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000203' }) MERGE (node)-[:usedBy { label:"Mid Suffolk"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000092' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000092' }) MERGE (node)-[:usedBy { label:"Rushmoor"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000029' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000029' }) MERGE (node)-[:usedBy { label:"Copeland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000054' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000054' }) MERGE (node)-[:usedBy { label:"Wiltshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000040' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000040' }) MERGE (node)-[:usedBy { label:"West Lothian"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000204' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000204' }) MERGE (node)-[:usedBy { label:"St Edmundsbury"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000152' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000152' }) MERGE (node)-[:usedBy { label:"East Northamptonshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000093' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000093' }) MERGE (node)-[:usedBy { label:"Test Valley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000030' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000030' }) MERGE (node)-[:usedBy { label:"Eden"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000031' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000031' }) MERGE (node)-[:usedBy { label:"Peterborough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000041' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000041' }) MERGE (node)-[:usedBy { label:"Angus"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000153' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000153' }) MERGE (node)-[:usedBy { label:"Kettering"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000205' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000205' }) MERGE (node)-[:usedBy { label:"Suffolk Coastal"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000031' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000031' }) MERGE (node)-[:usedBy { label:"South Lakeland"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000094' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000094' }) MERGE (node)-[:usedBy { label:"Winchester"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000032' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000032' }) MERGE (node)-[:usedBy { label:"Luton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000042' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000042' }) MERGE (node)-[:usedBy { label:"Dundee City"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000154' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000154' }) MERGE (node)-[:usedBy { label:"Northampton"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000206' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000206' }) MERGE (node)-[:usedBy { label:"Waveney"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000055' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000055' }) MERGE (node)-[:usedBy { label:"Bedford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000095' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000095' }) MERGE (node)-[:usedBy { label:"Broxbourne"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000032' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000032' }) MERGE (node)-[:usedBy { label:"Amber Valley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000044' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000044' }) MERGE (node)-[:usedBy { label:"North Lanarkshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000207' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000207' }) MERGE (node)-[:usedBy { label:"Elmbridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000155' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000155' }) MERGE (node)-[:usedBy { label:"South Northamptonshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000096' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000096' }) MERGE (node)-[:usedBy { label:"Dacorum"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000056' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000056' }) MERGE (node)-[:usedBy { label:"Central Bedfordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000033' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000033' }) MERGE (node)-[:usedBy { label:"Bolsover"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000045' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000045' }) MERGE (node)-[:usedBy { label:"East Dunbartonshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000156' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000156' }) MERGE (node)-[:usedBy { label:"Wellingborough"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000242' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000242' }) MERGE (node)-[:usedBy { label:"East Hertfordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E06000033' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E06000033' }) MERGE (node)-[:usedBy { label:"Southend-on-Sea"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000034' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000034' }) MERGE (node)-[:usedBy { label:"Chesterfield"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000208' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000208' }) MERGE (node)-[:usedBy { label:"Epsom and Ewell"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'S12000046' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'S12000046' }) MERGE (node)-[:usedBy { label:"Glasgow City"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000098' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000098' }) MERGE (node)-[:usedBy { label:"Hertsmere"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000035' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000035' }) MERGE (node)-[:usedBy { label:"Derbyshire Dales"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000209' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000209' }) MERGE (node)-[:usedBy { label:"Guildford"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000001' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000001' }) MERGE (node)-[:usedBy { label:"Antrim and Newtownabbey"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000099' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000099' }) MERGE (node)-[:usedBy { label:"North Hertfordshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000210' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000210' }) MERGE (node)-[:usedBy { label:"Mole Valley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000036' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000036' }) MERGE (node)-[:usedBy { label:"Erewash"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000002' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000002' }) MERGE (node)-[:usedBy { label:"Armagh City, Banbridge and Craigavon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000240' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000240' }) MERGE (node)-[:usedBy { label:"St Albans"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000037' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000037' }) MERGE (node)-[:usedBy { label:"High Peak"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000211' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000211' }) MERGE (node)-[:usedBy { label:"Reigate and Banstead"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000003' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000003' }) MERGE (node)-[:usedBy { label:"Belfast"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000004' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000004' }) MERGE (node)-[:usedBy { label:"Causeway Coast and Glens"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000212' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000212' }) MERGE (node)-[:usedBy { label:"Runnymede"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000005' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000005' }) MERGE (node)-[:usedBy { label:"Derry City and Strabane"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000213' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000213' }) MERGE (node)-[:usedBy { label:"Spelthorne"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000006' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000006' }) MERGE (node)-[:usedBy { label:"Fermanagh and Omagh"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000214' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000214' }) MERGE (node)-[:usedBy { label:"Surrey Heath"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000007' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000007' }) MERGE (node)-[:usedBy { label:"Lisburn and Castlereagh"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000215' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000215' }) MERGE (node)-[:usedBy { label:"Tandridge"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000008' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000008' }) MERGE (node)-[:usedBy { label:"Mid and East Antrim"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000216' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000216' }) MERGE (node)-[:usedBy { label:"Waverley"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000009' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000009' }) MERGE (node)-[:usedBy { label:"Mid Ulster"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000217' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000217' }) MERGE (node)-[:usedBy { label:"Woking"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000010' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000010' }) MERGE (node)-[:usedBy { label:"Newry, Mourne and Down"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000218' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000218' }) MERGE (node)-[:usedBy { label:"North Warwickshire"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'N09000011' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'N09000011' }) MERGE (node)-[:usedBy { label:"Ards and North Down"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000219' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000219' }) MERGE (node)-[:usedBy { label:"Nuneaton and Bedworth"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000220' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000220' }) MERGE (node)-[:usedBy { label:"Rugby"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000221' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000221' }) MERGE (node)-[:usedBy { label:"Stratford-on-Avon"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000222' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000222' }) MERGE (node)-[:usedBy { label:"Warwick"}]->(parent);

MERGE (node:`_code`:`_code_geography` { value:'E07000223' });

MATCH (parent:`_code_list`:`_code_list_local-authority`:`_geography`{geography_list_ID:'lad16'}),(node:`_code`:`_code_geography` { value:'E07000223' }) MERGE (node)-[:usedBy { label:"Adur"}]->(parent);