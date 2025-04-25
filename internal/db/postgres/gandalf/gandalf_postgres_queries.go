package gandalfpg

func (pdb *GandalfPostgres) fetchGetEntityDetailsQuery() string {
	return `
		SELECT 
			e.name AS club_name,
			e.founder AS founder,
			e.co_founder AS co_founder,
			e.club_email AS club_email,
			e.description AS club_description,
			e.year_established AS year_established,
			e.club_logo_image_url AS club_logo_image_url,
			e.club_banner_image_url AS club_banner_image_url,
			e.club_website_url AS club_website_url,
			e.club_twitter_url AS club_twitter_url,
			e.club_youtube_url AS club_youtube_url,
			e.club_facebook_url AS club_facebook_url,
			e.club_linkedin_url AS club_linkedin_url,
			e.club_instagram_url AS club_instagram_url,
			json_agg(json_build_object(
				'name', cm.name,
				'title', cm.title,
				'image_url', cm.image_url
			)) AS current_core_members
		FROM 
			entities e
		LEFT JOIN 
			present_core_members cm ON e.id = cm.eid
		GROUP BY 
			e.id
		ORDER BY 
			e.id;
	`
}

func (pdb *GandalfPostgres) fetchPostEntityDetailsQuery() string {
	return `
		INSERT INTO entities (
    		id,
    		name,
    		description,
    		year_established,
    		founder,
    		co_founder,
    		club_logo_image_url,
    		club_banner_image_url,
    		club_website_url,
    		club_facebook_url,
    		club_twitter_url,
    		club_instagram_url,
    		club_youtube_url,
    		club_linkedin_url,
    		club_email
		) VALUES (
    		$1,
    		$2,
    		$3,
    		$4,
    		$5,
    		$6,
    		$7,
    		$8,
    		$9,
    		$10,
    		$11,
    		$12,
    		$13,
    		$14,
    		$15
		) ON CONFLICT (id) DO UPDATE SET 
    		name = EXCLUDED.name,
    		description = EXCLUDED.description,
    		year_established = EXCLUDED.year_established,
    		founder = EXCLUDED.founder,
    		co_founder = EXCLUDED.co_founder,
    		club_logo_image_url = EXCLUDED.club_logo_image_url,
    		club_banner_image_url = EXCLUDED.club_banner_image_url,
    		club_website_url = EXCLUDED.club_website_url,
    		club_facebook_url = EXCLUDED.club_facebook_url,
    		club_twitter_url = EXCLUDED.club_twitter_url,
    		club_instagram_url = EXCLUDED.club_instagram_url,
    		club_youtube_url = EXCLUDED.club_youtube_url,
    		club_linkedin_url = EXCLUDED.club_linkedin_url,
    		club_email = EXCLUDED.club_email
		RETURNING id, name;
	`
}

func (pdb *GandalfPostgres) fetchPostCurrentCoreMembersQuery() string {
	return `
		INSERT INTO present_core_members (
		    eid,
		    name,
		    title,
		    image_url
		) VALUES (
		    $1,
		    $2,
		    $3,
		    $4
		);
	`
}

func (pdb *GandalfPostgres) fetchGetLegacyHoldersQuery() string {
	query := `
		SELECT
			jsonb_object_agg(year_range, legacy_holders_list ORDER BY year_range DESC) AS year_wise_legacy_holders
		FROM (
			SELECT
				year_start || '-' || year_end AS year_range,
				jsonb_agg(
					jsonb_build_object(
						'name', name,
						'title', title,
						'image_url', image_url
					) ORDER BY legacy_id
				) AS legacy_holders_list
			FROM
				legacy_holders
			WHERE
				eid = $1
			GROUP BY
				year_start, year_end
			ORDER BY
				year_start DESC, year_end DESC
		) AS yearly_aggregated_data;
	`
	return query
}
