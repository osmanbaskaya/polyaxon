// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.threeten.bp.OffsetDateTime;

/**
 * V1CodeReference
 */

public class V1CodeReference {
  @SerializedName("uuid")
  private String uuid = null;

  @SerializedName("commit")
  private String commit = null;

  @SerializedName("updated_at")
  private OffsetDateTime updatedAt = null;

  @SerializedName("status")
  private String status = null;

  @SerializedName("git_url")
  private String gitUrl = null;

  @SerializedName("is_dirty")
  private Boolean isDirty = null;

  public V1CodeReference uuid(String uuid) {
    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @ApiModelProperty(value = "")
  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  public V1CodeReference commit(String commit) {
    this.commit = commit;
    return this;
  }

   /**
   * Get commit
   * @return commit
  **/
  @ApiModelProperty(value = "")
  public String getCommit() {
    return commit;
  }

  public void setCommit(String commit) {
    this.commit = commit;
  }

  public V1CodeReference updatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }

  public void setUpdatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
  }

  public V1CodeReference status(String status) {
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @ApiModelProperty(value = "")
  public String getStatus() {
    return status;
  }

  public void setStatus(String status) {
    this.status = status;
  }

  public V1CodeReference gitUrl(String gitUrl) {
    this.gitUrl = gitUrl;
    return this;
  }

   /**
   * Get gitUrl
   * @return gitUrl
  **/
  @ApiModelProperty(value = "")
  public String getGitUrl() {
    return gitUrl;
  }

  public void setGitUrl(String gitUrl) {
    this.gitUrl = gitUrl;
  }

  public V1CodeReference isDirty(Boolean isDirty) {
    this.isDirty = isDirty;
    return this;
  }

   /**
   * Get isDirty
   * @return isDirty
  **/
  @ApiModelProperty(value = "")
  public Boolean isIsDirty() {
    return isDirty;
  }

  public void setIsDirty(Boolean isDirty) {
    this.isDirty = isDirty;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1CodeReference v1CodeReference = (V1CodeReference) o;
    return Objects.equals(this.uuid, v1CodeReference.uuid) &&
        Objects.equals(this.commit, v1CodeReference.commit) &&
        Objects.equals(this.updatedAt, v1CodeReference.updatedAt) &&
        Objects.equals(this.status, v1CodeReference.status) &&
        Objects.equals(this.gitUrl, v1CodeReference.gitUrl) &&
        Objects.equals(this.isDirty, v1CodeReference.isDirty);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, commit, updatedAt, status, gitUrl, isDirty);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1CodeReference {\n");
    
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    commit: ").append(toIndentedString(commit)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    gitUrl: ").append(toIndentedString(gitUrl)).append("\n");
    sb.append("    isDirty: ").append(toIndentedString(isDirty)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

